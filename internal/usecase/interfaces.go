// Package usecase implements application business logic. Each logic group in own file.
package usecase

import (
	"context"

	"gexec/internal/entity"
)

//go:generate mockgen -source=interfaces.go -destination=./mocks_test.go -package=usecase_test

type (
	// Login -.
	Exec interface {
		Save(context.Context, entity.User) error
		Query(context.Context, entity.User) error
	}

	// ExecRepo -.
	ExecRepo interface {
		Save(context.Context) error
		Query(context.Context) error
	}
)
