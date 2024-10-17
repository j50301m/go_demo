package kgserr

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

// ErrorInterceptor is a gRPC unary interceptor that converts custom KgsError to gRPC error with details
//
// Usage:
//
//	s := grpc.NewServer(
//		grpc.ChainUnaryInterceptor(
//			kgserr.ErrorInterceptor,
//			// Any other interceptors can be added here
//		),
func ErrorInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	resp, err := handler(ctx, req)
	if err == nil {
		return resp, nil
	}

	// Check if the error is our custom KgsError
	if kgsErr, ok := err.(*KgsError); ok {
		// Convert KgsError to gRPC status
		st := status.New(kgsErr.Code().GrpcCode(), kgsErr.Error())

		// Convert KgsError to proto
		proto, err := kgsErr.toProto()
		if err != nil {
			// If we can't add the details, just return the status as is
			return nil, st.Err()
		}
		// Add the details to the status
		detailedStatus, err := st.WithDetails(proto)
		if err != nil {
			// If we can't add the details, just return the status as is
			return nil, st.Err()
		}

		return nil, detailedStatus.Err()
	}

	// If it's not a KgsError, return the original error
	return nil, err
}

// StreamErrorInterceptor is a gRPC stream interceptor that converts custom KgsError to gRPC error with details
//
// Usage:
//
//	s := grpc.NewServer(
//		grpc.StreamInterceptor(
//			kgserr.StreamErrorInterceptor,
//			// Any other interceptors can be added here
//		),
func StreamErrorInterceptor(srv any, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	err := handler(srv, ss)
	if err == nil {
		return nil
	}

	// Check if the error is our custom KgsError
	if kgsErr, ok := err.(*KgsError); ok {
		// Convert KgsError to gRPC status
		st := status.New(kgsErr.Code().GrpcCode(), kgsErr.Error())

		// Convert KgsError to proto
		proto, err := kgsErr.toProto()
		if err != nil {
			// If we can't add the details, just return the status as is
			return st.Err()
		}
		// Add the details to the status
		detailedStatus, err := st.WithDetails(proto)
		if err != nil {
			// If we can't add the details, just return the status as is
			return st.Err()
		}

		return detailedStatus.Err()
	}

	// If it's not a KgsError, return the original error
	return err
}
