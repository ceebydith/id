/*
Package id provides utilities for generating and validating sequential identifiers using the Luhn algorithm or other methods or without validation.

The package includes the following components:
1. Sequencer: Interface for generating sequences.
2. Validator: Interface for signing and verifying numbers.
3. Generator: Struct for generating signed numbers with a sequence and validator.

Types:
  - Sequencer: An interface that defines the Generate method for creating sequences of int64.
  - Validator: An interface that defines the Sign and Verify methods for signing and verifying numbers.

Functions:
  - New(sequencer Sequencer, validator Validator, since ...int64) *Generator: Creates a new Generator instance with the given Sequencer, Validator, and optional start time.

Structs:
  - Generator: Generates signed numbers using a Sequencer and Validator.
    Methods:
  - (g *Generator) Generate(ctx context.Context) (int64, error): Produces a new signed number.
  - (g *Generator) Valid(value int64) bool: Verifies if the provided value is valid.

Example usage:

	package main

	import (
	    "context"
	    "fmt"
	    "myapp/id"
	)

	func main() {
	    sequencer := id.RangeSequencer(1, 100)
	    validator := id.LuhnValidator()
	    generator := id.New(sequencer, validator, time.Now().Unix())

	    num, err := generator.Generate(context.Background())
	    if err != nil {
	        fmt.Println("Error generating number:", err)
	    } else {
	        fmt.Println("Generated number:", num)
	    }

	    valid := generator.Valid(num)
	    fmt.Println("Is the number valid?", valid)
	}
*/
package id

import (
	"context"
	"time"
)

// Sequencer is an interface for generating sequence numbers
type Sequencer interface {
	Generate(ctx context.Context) (int64, error)
}

// Validator is an interface for signing and verifying numbers
type Validator interface {
	Sign(number int64) int64
	Verify(number int64) bool
}

// Generator generates unique IDs based on Unix time and a sequencer
type Generator struct {
	sequencer Sequencer
	validator Validator
	since     int64
}

// Generate creates a new unique ID by combining the current Unix time and a sequence number
func (g *Generator) Generate(ctx context.Context) (int64, error) {
	seq, err := g.sequencer.Generate(ctx)
	if err != nil {
		return 0, err
	}
	number := ((time.Now().Unix() - g.since) * 10000) + (seq % 10000)
	return g.validator.Sign(number), nil
}

// Valid verifies the given ID using the validator
func (g *Generator) Valid(value int64) bool {
	return g.validator.Verify(value)
}

// New creates a new Generator with the provided sequencer and validator.
// An optional start time can be specified.
func New(sequencer Sequencer, validator Validator, since ...time.Time) *Generator {
	startTime := int64(0)
	if len(since) > 0 {
		startTime = since[0].Unix()
	}
	return &Generator{
		sequencer: sequencer,
		validator: validator,
		since:     startTime,
	}
}
