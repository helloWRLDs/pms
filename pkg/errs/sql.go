package errs

import (
	"database/sql"
	"errors"
	"fmt"
)

type RepositoryDetails struct {
	Operation string
	Object    string
	Field     string
	Value     string
}

func WithOperation(operation string) func(*RepositoryDetails) {
	return func(rd *RepositoryDetails) {
		rd.Operation = operation
	}
}

func WithField(field, value string) func(*RepositoryDetails) {
	return func(rd *RepositoryDetails) {
		rd.Field = field
		rd.Value = value
	}
}

func WithObject(object string) func(*RepositoryDetails) {
	return func(rd *RepositoryDetails) {
		rd.Object = object
	}
}

func (rd RepositoryDetails) MapSQL(err error, opts ...func(*RepositoryDetails)) error {
	if err == nil {
		return nil
	}
	for _, fn := range opts {
		fn(&rd)
	}
	switch {
	case errors.Is(err, sql.ErrNoRows):
		return ErrNotFound{
			Object: rd.Object,
			Field:  rd.Field,
			Value:  rd.Value,
		}
	default:
		return ErrInternal{
			Reason: fmt.Sprintf("failed to %s %s", rd.Operation, rd.Object),
		}
	}
}
