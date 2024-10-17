package responder

import (
	"context"
	"encoding/json"
	"errors"
	"hype-casino-platform/pkg/kgserr"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/otel/trace"
)

// TestOk verifies that the Ok function correctly creates a success response.
// It checks if the response code, message, and data are set as expected.
func TestOk(t *testing.T) {
	data := map[string]string{"key": "value"}
	response := Ok(data)

	assert.Equal(t, kgserr.OK, int(response.code))
	assert.Equal(t, data, response.data)
}

// TestError ensures that the Error function properly handles KgsError.
// It verifies if the response code, error message, and data are set correctly.
func TestError(t *testing.T) {
	err := kgserr.New(kgserr.InvalidArgument, "Invalid input", nil)
	response := Error(err)

	assert.Equal(t, kgserr.InvalidArgument, int(response.code))
	assert.Nil(t, response.data)
}

// TestUnknownError checks if the UnknownError function correctly handles
// generic errors by setting the appropriate error code and message.
func TestUnknownError(t *testing.T) {
	err := errors.New("Unknown error occurred")
	response := UnknownError(err)

	assert.Equal(t, kgserr.InternalServerError, int(response.code))
	assert.Nil(t, response.data)
}

// TestResponseToGinH verifies that the toGinH method correctly converts
// a Response struct to a gin.H map with the expected key-value pairs.
func TestResponseToGinH(t *testing.T) {
	response := &Response{
		code: kgserr.OK,
		data: map[string]string{"key": "value"},
	}

	ginH := response.toGinH("traceId")
	assert.Equal(t, kgserr.KgsCode(200_0000), ginH["code"])
	assert.Equal(t, map[string]string{"key": "value"}, ginH["data"])
}

// TestResponseHttpCode ensures that the HttpCode method returns
// the correct HTTP status code based on the response's KgsCode.
func TestResponseHttpCode(t *testing.T) {
	response := &Response{code: kgserr.BadRequest}
	assert.Equal(t, http.StatusBadRequest, response.HttpCode())
}

// TestResponseWithContext verifies that the WithContext method
// correctly stores the Response in the Gin context.
func TestResponseWithContext(t *testing.T) {
	gin.SetMode(gin.TestMode)
	c, _ := gin.CreateTestContext(nil)

	response := Ok(nil)
	response.WithContext(c)

	value, exists := c.Get(_responseKey)
	require.True(t, exists)
	assert.Equal(t, response, value)
}

func createTestContext(traceID trace.TraceID) (*gin.Context, *httptest.ResponseRecorder) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// Create a span context with the fake TraceID
	sc := trace.NewSpanContext(trace.SpanContextConfig{
		TraceID: traceID,
		SpanID:  trace.SpanID{},
		Remote:  false,
	})

	// Set the span context in the request context
	ctx := trace.ContextWithSpanContext(context.Background(), sc)
	c.Request = httptest.NewRequest(http.MethodGet, "/", nil).WithContext(ctx)

	return c, w
}

// TestGinResponser tests the GinResponser middleware under various scenarios,
// including handling KgsErrors, unknown errors, successful responses, and no responses.
func TestGinResponser(t *testing.T) {
	gin.SetMode(gin.TestMode)

	var response map[string]interface{}

	// Create a fake TraceID
	fakeTraceID := trace.TraceID{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0A, 0x0B, 0x0C, 0x0D, 0x0E, 0x0F}

	// Test handling of KgsError
	t.Run("Handle KgsError", func(t *testing.T) {
		c, w := createTestContext(fakeTraceID)
		_ = c.Error(kgserr.New(kgserr.BadRequest, "Bad Request", nil))

		middleware := GinResponser()
		middleware(c)

		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)

		assert.Equal(t, fakeTraceID.String(), response["traceId"])
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	// Test handling of unknown error
	t.Run("Handle Unknown Error", func(t *testing.T) {
		c, w := createTestContext(fakeTraceID)
		_ = c.Error(errors.New("Unknown Error"))

		middleware := GinResponser()
		middleware(c)

		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)

		assert.Equal(t, fakeTraceID.String(), response["traceId"])
		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})

	// Test handling of successful response
	t.Run("Handle Response", func(t *testing.T) {
		c, w := createTestContext(fakeTraceID)
		Ok(nil).WithContext(c)

		middleware := GinResponser()
		middleware(c)

		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)

		assert.Equal(t, fakeTraceID.String(), response["traceId"])
		assert.Equal(t, http.StatusOK, w.Code)
	})

	// Test handling of no response
	t.Run("Handle No Response", func(t *testing.T) {
		c, w := createTestContext(fakeTraceID)

		middleware := GinResponser()
		middleware(c)

		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)

		assert.Equal(t, fakeTraceID.String(), response["traceId"])
		assert.Equal(t, http.StatusNotFound, w.Code)
	})
}
