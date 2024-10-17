package kgserr

import (
	"encoding/json"
	"errors"
	"fmt"
	internal "hype-casino-platform/pkg/kgserr/internal/gen"
	"net/http"

	"google.golang.org/grpc/status"
)

// KgsError is a custom error type that wraps error code, message, and error sources.
type KgsError struct {
	code    KgsCode
	msg     string
	data    any
	sources []error
}

// New creates a new KgsError.
// Parameters:
//   - code: The KgsCode of the error.
//   - msg: The error message.
//   - data: The data associated with the error.
//   - source: The error sources.
//
// Returns:
//   - error: The KgsError.
//
// Example:
//
//	err := NewKgsError(ErrInvalidInput, "err_msg")
//	err := NewKgsError(ErrInvalidInput, "err_msg",err)
//	err := NewKgsError(ErrInvalidInput, "err_msg",err1,err2)
func New(code KgsCode, msg string, source ...error) *KgsError {
	return &KgsError{
		code:    code,
		msg:     msg,
		sources: source,
	}
}

// Error returns a string representation of the KgsError.
// if there are sources, it will include the sources in the string.
func (e *KgsError) Error() string {
	if len(e.sources) > 0 {
		return fmt.Sprintf("kgsCode: %v, msg:%s,  sources: %v", e.code, e.msg, e.sources)
	}
	return fmt.Sprintf("kgsCode: %v, msg:%s", e.code, e.msg)
}

// HttpCode returns the standard HTTP status code.
func (e *KgsError) HttpCode() int {
	// Check self is nil
	if e == nil {
		return http.StatusInternalServerError
	}

	return e.code.HttpCode()
}

// Code returns the kgs error code.
func (e *KgsError) Code() KgsCode {
	return e.code
}

// Message returns the error message.
func (e *KgsError) Message() string {
	return e.msg
}

// Unwrap returns the error sources.
// if there are no sources, it will return nil.
func (e *KgsError) Unwrap() []error {
	return e.sources
}

// WithData add data to the KgsError.
// Parameters:
//   - data: The data to be added to the KgsError.
//
// Returns:
//   - error: The KgsError.
//
// Example:
//
//	data := map[string]interface{}{"key1": "value1"}
//	err := NewKgsError(ErrInvalidInput, "err_msg").WithData(data)
func (e *KgsError) WithData(data any) *KgsError {
	e.data = data
	return e
}

// Is checks if the target error matches the KgsError.
func (e *KgsError) Is(target error) bool {
	t, ok := target.(*KgsError)
	if !ok {
		return false
	}
	return e.code == t.code
}

// Data returns the data associated with the KgsError.
func (e *KgsError) Data() any {
	return e.data
}

// WithSource add error sources to the KgsError.
func (e *KgsError) WithSource(err error) *KgsError {
	e.sources = append(e.sources, err)
	return e
}

// FromGrpcErr converts a gRPC error to a KgsError.
// Parameters:
//   - err: The gRPC error.
//
// Returns:
//   - error: The KgsError.
//   - ok: A boolean indicating if the conversion was successful.
//
// Example:
//
//	kgsErr, ok := FromGrpcErr(err)
func FromGrpcErr(err error) (kgsErr *KgsError, ok bool) {
	st, ok := status.FromError(err)
	if !ok {
		return nil, false
	}

	// Check if the error is our custom KgsError
	for _, detail := range st.Details() {
		if proto, ok := detail.(*internal.KgsErrorProto); ok {
			kgsErr, err := fromProto(proto)
			if err != nil {
				return nil, false
			}
			return kgsErr, true
		}
	}

	return nil, false
}

// toProto converts the KgsError to a proto message.
func (e *KgsError) toProto() (*internal.KgsErrorProto, error) {
	dataBytes, err := json.Marshal(e.data)
	if err != nil {
		return nil, err
	}

	sources := make([]string, len(e.sources))
	for i, src := range e.sources {
		if src != nil {
			sources[i] = src.Error()
		}
	}

	return &internal.KgsErrorProto{
		Code:    int32(e.code),
		Message: e.msg,
		Data:    dataBytes,
		Source:  sources,
	}, nil
}

// fromProto converts a proto message to a KgsError.
func fromProto(proto *internal.KgsErrorProto) (*KgsError, error) {
	data := make(map[string]interface{})
	if err := json.Unmarshal(proto.Data, &data); err != nil {
		return nil, err
	}

	sources := make([]error, len(proto.Source))
	for i, src := range proto.Source {
		sources[i] = errors.New(src)
	}

	return &KgsError{
		code:    KgsCode(proto.Code),
		msg:     proto.Message,
		data:    data,
		sources: sources,
	}, nil
}
