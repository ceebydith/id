package id_test

import (
	"context"
	"testing"
	"time"

	"github.com/ceebydith/id"
	"github.com/stretchr/testify/assert"
)

func TestRangeSequencer(t *testing.T) {
	seq := id.RangeSequencer(1, 3)

	val, err := seq.Generate(context.TODO())
	assert.NoError(t, err)
	assert.Equal(t, int64(1), val)

	val, err = seq.Generate(context.TODO())
	assert.NoError(t, err)
	assert.Equal(t, int64(2), val)

	val, err = seq.Generate(context.TODO())
	assert.NoError(t, err)
	assert.Equal(t, int64(3), val)

	val, err = seq.Generate(context.TODO())
	assert.NoError(t, err)
	assert.Equal(t, int64(1), val)
}

func TestLuhnValidator(t *testing.T) {
	val := id.LuhnValidator()

	signed := val.Sign(123456789)
	assert.True(t, val.Verify(signed))

	assert.False(t, val.Verify(1234567890))
}

func TestGenerator(t *testing.T) {
	seq := id.RangeSequencer(1, 100)
	val := id.LuhnValidator()
	gen := id.New(seq, val, time.Now())

	num, err := gen.Generate(context.TODO())
	assert.NoError(t, err)
	assert.True(t, gen.Valid(num))

	assert.False(t, gen.Valid(num+1))
}

func TestWithoutValidator(t *testing.T) {
	val := id.NoValidator()

	signed := val.Sign(123456789)
	assert.True(t, val.Verify(signed))

	assert.True(t, val.Verify(1234567890))
}
