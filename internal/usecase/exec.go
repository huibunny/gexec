package usecase

import (
	"context"

	"gexec/internal/entity"
	"gexec/internal/usecase/repo"
)

// ExecUserCase -.
type ExecUserCase struct {
	repo *repo.UserRepo
}

// New -.
func New(r *repo.UserRepo) *ExecUserCase {
	return &ExecUserCase{
		repo: r,
	}
}

// Query -.
func (uc *ExecUserCase) Save(ctx context.Context, t entity.User) error {

	return nil
}

// Query -.
func (uc *ExecUserCase) Query(ctx context.Context, t entity.User) error {

	return nil
}
