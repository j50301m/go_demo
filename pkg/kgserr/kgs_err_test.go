package kgserr

import (
	"errors"
	internal "hype-casino-platform/pkg/kgserr/internal/gen"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// TestKgsCodeHttpCode tests the HttpCode method of KgsCode
func TestKgsCodeHttpCode(t *testing.T) {
	tests := []struct {
		name     string
		code     KgsCode
		expected int
	}{
		{"OK", OK, http.StatusOK},
		{"BadRequest", BadRequest, http.StatusBadRequest},
		{"Unauthorized", Unauthorized, http.StatusUnauthorized},
		{"StatusNotFound", ResponseNotFound, http.StatusNotFound},
		{"InternalServerError", InternalServerError, http.StatusInternalServerError},
		{"NotImplemented", NotImplemented, http.StatusNotImplemented},
		{"InvalidCode", 999_9999, http.StatusInternalServerError},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.code.HttpCode())
		})
	}
}

// TestNewKgsError tests the New function for creating KgsError
func TestNewKgsError(t *testing.T) {
	err := New(BadRequest, "test error")
	assert.Equal(t, BadRequest, err.Code().Int())
	assert.Equal(t, "test error", err.Message())
	assert.Nil(t, err.Data())
	assert.Empty(t, err.Unwrap())

	sourceErr := errors.New("source error")
	err = New(BadRequest, "test error", sourceErr)
	assert.Equal(t, BadRequest, err.Code().Int())
	assert.Equal(t, "test error", err.Message())
	assert.Nil(t, err.Data())
	assert.Equal(t, []error{sourceErr}, err.Unwrap())
}

// TestKgsErrorError tests the Error method of KgsError
func TestKgsErrorError(t *testing.T) {
	err := New(BadRequest, "test error")
	assert.Equal(t, "kgsCode: 4000000, msg:test error", err.Error())

	sourceErr := errors.New("source error")
	err = New(BadRequest, "test error", sourceErr)
	assert.Equal(t, "kgsCode: 4000000, msg:test error,  sources: [source error]", err.Error())
}

// TestKgsErrorHttpCode tests the HttpCode method of KgsError
func TestKgsErrorHttpCode(t *testing.T) {
	err := New(BadRequest, "test error")
	assert.Equal(t, http.StatusBadRequest, err.HttpCode())

	var nilErr *KgsError
	assert.Equal(t, http.StatusInternalServerError, nilErr.HttpCode())
}

// TestKgsErrorWithData tests the WithData method of KgsError
func TestKgsErrorWithData(t *testing.T) {
	data := map[string]string{"key": "value"}
	err := New(BadRequest, "test error").WithData(data)
	assert.Equal(t, data, err.Data())
}

// TestKgsErrorIs tests the Is method of KgsError
func TestKgsErrorIs(t *testing.T) {
	err1 := New(BadRequest, "test error")
	err2 := New(BadRequest, "another test error")
	err3 := New(Unauthorized, "unauthorized error")

	assert.True(t, err1.Is(err2))
	assert.False(t, err1.Is(err3))
	assert.False(t, err1.Is(errors.New("standard error")))
}

// TestKgsErrorWithSource tests the WithSource method of KgsError
func TestKgsErrorWithSource(t *testing.T) {
	sourceErr := errors.New("source error")
	err := New(BadRequest, "test error").WithSource(sourceErr)
	assert.Equal(t, []error{sourceErr}, err.Unwrap())
}

func TestFromGrpcErr(t *testing.T) {
	tests := []struct {
		name    string
		err     error
		wantErr *KgsError
		wantOk  bool
	}{
		{
			name:    "Non-gRPC error",
			err:     errors.New("regular error"),
			wantErr: nil,
			wantOk:  false,
		},
		{
			name:    "gRPC error without KgsErrorProto",
			err:     status.Error(codes.NotFound, "not found"),
			wantErr: nil,
			wantOk:  false,
		},
		{
			name: "gRPC error with KgsErrorProto",
			err: func() error {
				st := status.New(codes.Internal, "internal error")
				kgsProto, _ := (&KgsError{
					code:    500,
					msg:     "test error",
					data:    map[string]interface{}{"key": "value"},
					sources: []error{errors.New("source error")},
				}).toProto()
				st, _ = st.WithDetails(kgsProto)
				return st.Err()
			}(),
			wantErr: &KgsError{
				code:    500,
				msg:     "test error",
				data:    map[string]interface{}{"key": "value"},
				sources: []error{errors.New("source error")},
			},
			wantOk: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotErr, gotOk := FromGrpcErr(tt.err)
			assert.Equal(t, tt.wantOk, gotOk)
			if tt.wantErr != nil {
				assert.Equal(t, tt.wantErr.code, gotErr.code)
				assert.Equal(t, tt.wantErr.msg, gotErr.msg)
				assert.Equal(t, tt.wantErr.data, gotErr.data)
				assert.Equal(t, len(tt.wantErr.sources), len(gotErr.sources))
				for i := range tt.wantErr.sources {
					assert.Equal(t, tt.wantErr.sources[i].Error(), gotErr.sources[i].Error())
				}
			} else {
				assert.Nil(t, gotErr)
			}
		})
	}
}

func TestKgsErrorToProto(t *testing.T) {
	t.Run("Normal KgsError", func(t *testing.T) {
		kgsErr := &KgsError{
			code:    400,
			msg:     "bad request",
			data:    map[string]interface{}{"reason": "invalid input"},
			sources: []error{errors.New("validation failed")},
		}

		proto, err := kgsErr.toProto()
		assert.NoError(t, err)
		assert.NotNil(t, proto)

		assert.Equal(t, int32(400), proto.Code)
		assert.Equal(t, "bad request", proto.Message)
		assert.JSONEq(t, `{"reason":"invalid input"}`, string(proto.Data))
		assert.Equal(t, []string{"validation failed"}, proto.Source)
	})

	t.Run("nil data and sources", func(t *testing.T) {
		kgsErr := &KgsError{
			code: 500,
			msg:  "internal error",
		}

		proto, err := kgsErr.toProto()
		assert.NoError(t, err)
		assert.NotNil(t, proto)

		assert.Equal(t, int32(500), proto.Code)
		assert.Equal(t, "internal error", proto.Message)
		assert.Empty(t, proto.Source)
	})
}

func TestFromProto(t *testing.T) {
	t.Run("Normal KgsErrorProto", func(t *testing.T) {
		proto := &internal.KgsErrorProto{
			Code:    int32(404),
			Message: "not found",
			Data:    []byte(`{"id":"123"}`),
			Source:  []string{"database error"},
		}

		kgsErr, err := fromProto(proto)
		assert.NoError(t, err)
		assert.NotNil(t, kgsErr)

		assert.Equal(t, KgsCode(404), kgsErr.code)
		assert.Equal(t, "not found", kgsErr.msg)
		assert.Equal(t, map[string]interface{}{"id": "123"}, kgsErr.data)
		assert.Equal(t, 1, len(kgsErr.sources))
		assert.Equal(t, "database error", kgsErr.sources[0].Error())
	})

	t.Run("Invalid JSON data", func(t *testing.T) {
		proto := &internal.KgsErrorProto{
			Code:    int32(404),
			Message: "not found",
			Data:    []byte(`invalid json`),
			Source:  []string{"database error"},
		}

		kgsErr, err := fromProto(proto)
		assert.Error(t, err)
		assert.Nil(t, kgsErr)
	})

	t.Run("Empty source", func(t *testing.T) {
		proto := &internal.KgsErrorProto{
			Code:    int32(404),
			Message: "not found",
			Data:    []byte(`{"id":"123"}`),
		}

		kgsErr, err := fromProto(proto)
		assert.NoError(t, err)
		assert.NotNil(t, kgsErr)

		assert.Equal(t, KgsCode(404), kgsErr.code)
		assert.Equal(t, "not found", kgsErr.msg)
		assert.Equal(t, map[string]interface{}{"id": "123"}, kgsErr.data)
		assert.Empty(t, kgsErr.sources)
	})
}
