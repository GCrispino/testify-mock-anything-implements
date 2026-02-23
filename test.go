package test

import (
	"context"
)

type Caller interface {
	Call(context.Context) error
}

func MustCall(ctx context.Context, c Caller) {
	if err := c.Call(ctx); err != nil {
		panic(err)
	}
}
