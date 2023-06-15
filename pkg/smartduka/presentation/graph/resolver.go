package graph

import (
	"context"
	"log"

	"github.com/oryx-systems/smartduka/pkg/smartduka/usecases"
)

//go:generate go run github.com/99designs/gqlgen

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	smartduka *usecases.Smartduka
}

// NewResolver initializes a new resolver
func NewResolver(ctx context.Context, smartduka usecases.Smartduka) (*Resolver, error) {
	return &Resolver{
		smartduka: &smartduka,
	}, nil
}

func (r Resolver) checkPreconditions() {
	if r.smartduka == nil {
		log.Panicf("expected smartduka usecases to be defined resolver")
	}
}
