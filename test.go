package test

import (
	"context"
)

// Caller is a simple interface with a Call method that taksa a context and returns an error.
type Caller interface {
	Call(context.Context) error
}

// MustCall calls the Call method of the provided Caller with the given context.
func MustCall(ctx context.Context, c Caller) {
	if err := c.Call(ctx); err != nil {
		panic(err)
	}
}
