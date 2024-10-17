package responder

import (
	"hype-casino-platform/pkg/kgserr"

	"github.com/gin-gonic/gin"
)

// Response represents a standardized API response structure.
// It encapsulates the response code, message, and data.
type Response struct {
	code kgserr.KgsCode
	data any
}

// Ok creates a new Response instance for successful operations.
//
// Parameters:
//   - data: The data to be included in the response.
//
// Returns:
//   - A pointer to a Response struct with OK status and the provided data.
//
// Usage:
//
//	res := responder.Ok(someData).withContext(c)
func Ok(data any) *Response {
	return &Response{
		code: kgserr.OK,
		data: data,
	}
}

// Error creates a new Response instance for error scenarios.
//
// Parameters:
//   - err: A pointer to a KgsError containing error details.
//
// Returns:
//   - A pointer to a Response struct with error information.
//
// Usage:
//
//	res := responder.Error(kgserr.New(kgserr.NotFound, "Resource not found", nil))
func Error(err *kgserr.KgsError) *Response {
	return &Response{
		code: err.Code(),
		data: err.Data(),
	}
}

// UnknownError creates a new Response instance for unknown error scenarios.
// This method is used when an error occurs that is not a KgsError.
//
// Parameters:
//   - err: A standard error interface.
//
// Returns:
//   - A pointer to a Response struct with InternalServerError status and error message.
//
// Usage:
//
//	if err,ok:= err.(*kgserr.KgsError); !ok {
//			res := responder.UnknownError(err).withContext(c)
//	}
func UnknownError(err error) *Response {
	return &Response{
		code: kgserr.InternalServerError,
		data: nil,
	}
}

// toGinH converts the Response to a Gin H map for JSON serialization.
//
// Parameters:
//   - traceId: A string representing the trace ID for the
//     current request.
//
// Returns:
//   - A Gin H map representing the Response object.
func (r *Response) toGinH(traceId string) gin.H {
	return gin.H{
		"code":    r.code,
		"traceId": traceId,
		"data":    r.data,
	}
}

// HttpCode returns the HTTP status code associated with the Response.
//
// Returns:
//   - An integer representing the HTTP status code.
func (r *Response) HttpCode() int {
	return r.code.HttpCode()
}

// WithContext sets the Response in the given Gin context.
//
// This method allows storing the Response object in the Gin context
// for later retrieval and processing by middleware or other handlers.
//
// Parameters:
//   - c: A pointer to the gin.Context in which to store the Response.
//
// Usage:
//
//	res := responder.Ok(someData)
//	res.WithContext(c)
func (r *Response) WithContext(c *gin.Context) {
	c.Set(_responseKey, r)
}
