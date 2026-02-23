package test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/mock"
)

type MockCaller struct {
	mock.Mock
}

func (c *MockCaller) Call(ctx context.Context) error {
	args := c.Called(ctx)
	return args.Error(0)
}

type testInterface interface {
	X() string
}

// passes ✅, as context.Background and context.TODO both implement context.Context
func TestMustCallAnythingImplementing(t *testing.T) {
	m := MockCaller{}

	// try with context.Background
	ctx := context.Background()

	m.On("Call", mock.AnythingImplementing((*context.Context)(nil))).
		Return(nil)

	MustCall(ctx, &m)

	m.AssertExpectations(t)

	// try with context.TODO
	ctx = context.TODO()
	m.On("Call", mock.AnythingImplementing((*context.Context)(nil))).
		Return(nil)

	MustCall(ctx, &m)

	m.AssertExpectations(t)
}

// fails ❌, as context.Background doesn't implement testInterface
func TestMustCallAnythingImplementingFails(t *testing.T) {
	m := MockCaller{}
	ctx := context.Background()

	m.On("Call", mock.AnythingImplementing((*testInterface)(nil))).
		Return(nil)

	MustCall(ctx, &m)

	m.AssertExpectations(t)
}
