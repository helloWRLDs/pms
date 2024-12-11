package errs

import "fmt"

type errHTTP struct {
	Status  int    `json:"status"`
	Message string `json:"msg"`
	Err     string `json:"err"`
}

func (e errHTTP) Error() string {
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
}

func WrapHttp(err error) error {
	resultErr := errHTTP{
		Message: err.Error(),
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
	default:
		resultErr.Status = 500
	}
	resultErr.Err = httpMessages[resultErr.Status]

	return resultErr
}
