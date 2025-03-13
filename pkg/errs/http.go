package errs

import (
	"fmt"
	"net/http"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ErrHTTP struct {
	Status  int    `json:"status"`
	Message string `json:"msg"`
	Err     string `json:"err"`
}

func (e ErrHTTP) Error() string {
	return fmt.Sprintf("[%d] %s - %s", e.Status, e.Err, e.Message)
}

var httpMessages = map[int]string{
	400: "Bad Gateway",
	401: "Unauthorized",
	403: "Forbidden",
	404: "Not Found",
	409: "Conflict",
	422: "Unporcessable Entity",
	429: "Too Many Requests",
	500: "Internal Server Error",
	503: "Service Unavailable",
}

func GRPCtoHTTP(err error) error {
	var resultErr ErrHTTP

	st, ok := status.FromError(err)
	if !ok {
		resultErr.Status = http.StatusInternalServerError
		resultErr.Err = httpMessages[resultErr.Status]
		resultErr.Message = "Unknown error"
		return resultErr
	}

	switch st.Code() {
	case codes.InvalidArgument:
		resultErr.Status = http.StatusBadRequest
	case codes.Unauthenticated:
		resultErr.Status = http.StatusUnauthorized
	case codes.PermissionDenied:
		resultErr.Status = http.StatusForbidden
	case codes.NotFound:
		resultErr.Status = http.StatusNotFound
	case codes.AlreadyExists:
		resultErr.Status = http.StatusConflict
	case codes.ResourceExhausted:
		resultErr.Status = http.StatusTooManyRequests
	case codes.Internal:
		resultErr.Status = http.StatusInternalServerError
	case codes.Unavailable:
		resultErr.Status = http.StatusBadGateway
	case codes.DeadlineExceeded:
		resultErr.Status = http.StatusGatewayTimeout
	default:
		resultErr.Status = http.StatusInternalServerError
	}
	resultErr.Err = httpMessages[resultErr.Status]
	resultErr.Message = st.Message()

	return resultErr
}

func WrapHttp(err error) error {
	if err == nil {
		return nil
	}
	resultErr := ErrHTTP{
		Message: err.Error(),
	}
	if strings.HasPrefix(err.Error(), "rpc error") {
		return GRPCtoHTTP(err)
	}

	switch err.(type) {
	case ErrBadGateway:
		resultErr.Status = 400
	case ErrUnauthorized:
		resultErr.Status = 401
	case ErrForbidden:
		resultErr.Status = 403
	case ErrNotFound:
		resultErr.Status = 404
	case ErrConflict:
		resultErr.Status = 409
	case ErrInvalidInput:
		resultErr.Status = 422
	case ErrTooManyRequests:
		resultErr.Status = 429
	case ErrUnavalaiable:
		resultErr.Status = 503
	default:
		resultErr.Status = 500
	}
	resultErr.Err = httpMessages[resultErr.Status]

	return resultErr
}
