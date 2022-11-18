package main

import (
	"context"
)

type token struct{}

// Pool is used to limit a set of workers to a maximum size.
type Pool struct {
	tokens chan token
}

// NewPool returns a new work pool of the given size.
func NewPool(size int) *Pool {
	return &Pool{
		tokens: make(chan token, size),
	}
}

// Acquire returns when a work token is retrieved from the pool (when the
// worker is permitted to start).
//
// It returns ctx.Err() if the context is cancelled.
func (p *Pool) Acquire(ctx context.Context) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	case p.tokens <- token{}:
		return nil
	}
}

// Release returns a work token to the pool.
func (p *Pool) Release() {
	<-p.tokens
}

// Size returns the size of the pool
func (p *Pool) Size() int {
	return cap(p.tokens)
}
