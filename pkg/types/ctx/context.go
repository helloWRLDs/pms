package ctx

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"
)

type Context struct {
	context.Context
	tx *sqlx.Tx
}

func New(ctx context.Context) Context {
	return Context{
		Context: ctx,
		tx:      nil,
	}
}

func (c *Context) Deadline() (deadline time.Time, ok bool) {
	return c.Context.Deadline()
}

func (c *Context) Done() <-chan struct{} {
	return c.Context.Done()
}

func (c *Context) Err() error {
	return c.Context.Err()
}

func (c *Context) Value(key any) any {
	return c.Context.Value(key)
}
