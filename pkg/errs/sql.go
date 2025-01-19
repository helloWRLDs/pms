package errs

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"
)

type RepositoryDetails struct {
	DBType    string
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
	if errors.Is(err, sql.ErrNoRows) {
		return ErrNotFound{
			Object: rd.Object,
			Field:  rd.Field,
			Value:  rd.Value,
		}
	}
	if rd.DBType == "SQLITE" {
		if strings.Contains(err.Error(), "2067") {
			return ErrConflict{
				Reason: fmt.Sprintf("%s with %s = %s already exists", rd.Object, rd.Field, rd.Value),
			}
		}
	}

	return ErrInternal{
		Reason: fmt.Sprintf("failed to %s %s", rd.Operation, rd.Object),
	}
}
