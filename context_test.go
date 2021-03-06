package fluc

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

const (
	testKey1   = "k1"
	testValue1 = "v1"
	testKey2   = "k2"
	testValue2 = "v2"
	testKey3   = "k3"
)

func TestEmptyConstructor(t *testing.T) {
	fluentCtx := Context()
	assert.IsType(t, &Wrapper{}, fluentCtx)
}

func TestConstructorWithContext(t *testing.T) {
	fluentCtx := Context(context.Background())
	assert.IsType(t, &Wrapper{}, fluentCtx)
}

func TestConstructorWithInvalidArgs(t *testing.T) {
	assert.Panics(t, func() {
		_ = Context(nil)
	})

	assert.Panics(t, func() {
		_ = Context(context.Background(), context.Background())
	})
}

func TestGetter(t *testing.T) {
	fluentCtx := Context(context.Background())
	ctx := fluentCtx.Get()
	assert.Implements(t, (*context.Context)(nil), ctx)
}

func TestFluentValueSetter(t *testing.T) {
	fluentCtx := Context(context.Background()).WithValue(testKey1, testValue1).With(testKey2, testValue2)
	assert.IsType(t, &Wrapper{}, fluentCtx)

	ctx := fluentCtx.Get()
	assert.Exactly(t, testValue1, ctx.Value(testKey1))
	assert.Exactly(t, testValue2, ctx.Value(testKey2))
	assert.Nil(t, ctx.Value(testKey3))
}

func TestWithDeadlineGetter(t *testing.T) {
	fluentCtx := Context(context.Background())
	ctx, cancel := fluentCtx.WithDeadline(time.Now())
	assert.Implements(t, (*context.Context)(nil), ctx)
	assert.IsType(t, context.CancelFunc(func() {}), cancel)
}

func TestWithCancelGetter(t *testing.T) {
	fluentCtx := Context(context.Background())
	ctx, cancel := fluentCtx.WithCancel()
	assert.Implements(t, (*context.Context)(nil), ctx)
	assert.IsType(t, context.CancelFunc(func() {}), cancel)
}

func TestWithTimeoutGetter(t *testing.T) {
	fluentCtx := Context(context.Background())
	ctx, cancel := fluentCtx.WithTimeout(time.Second)
	assert.Implements(t, (*context.Context)(nil), ctx)
	assert.IsType(t, context.CancelFunc(func() {}), cancel)
}
