package errs

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func WrapGRPC(err error) error {
	if err == nil {
		return nil
	}

	var (
		resultErr error
		msg       = err.Error()
	)

	switch err.(type) {
	case ErrBadGateway, ErrInvalidInput:
		resultErr = status.Error(codes.InvalidArgument, msg)
	case ErrUnauthorized:
		resultErr = status.Error(codes.Unauthenticated, msg)
	case ErrForbidden:
		resultErr = status.Error(codes.PermissionDenied, msg)
	case ErrNotFound:
		resultErr = status.Error(codes.NotFound, msg)
	case ErrConflict, ErrAlreadyExist:
		resultErr = status.Error(codes.AlreadyExists, msg)
	case ErrTooManyRequests:
		resultErr = status.Error(codes.ResourceExhausted, msg)
	default:
		resultErr = status.Error(codes.Internal, msg)
	}
	return resultErr
}
