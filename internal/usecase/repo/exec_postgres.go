package repo

import (
	"context"
	"fmt"

	"gexec/pkg/postgres"
)

const _defaultEntityCap = 64

// UserRepo -.
type UserRepo struct {
	*postgres.Postgres
}

// New -.
func New(pg *postgres.Postgres) *UserRepo {
	return &UserRepo{pg}
}

// Save -.
func (r *UserRepo) Save(ctx context.Context, table string, columns []string, values []interface{}) (string, int, error) {
	sql, args, err := r.Builder.Insert(table).Columns(columns...).Values(values...).Suffix("RETURNING id").ToSql()
	if err != nil {
		return "", -1, fmt.Errorf("ExecRepo - Save - r.Builder: %w", err)
	}

	tx, err := r.Pool.Begin(ctx)
	if err != nil {
		return "", -3, fmt.Errorf("ExecRepo - Save - r.Pool.Begin: %w", err)
	}
	idList := ""
	err = tx.QueryRow(ctx, sql, args...).Scan(&idList)
	if err != nil {
		tx.Rollback(ctx)
		return "", -2, fmt.Errorf("ExecRepo - Save r.Pool.Query: %w", err)
	}

	err = tx.Commit(ctx)
	if err != nil {
		tx.Rollback(ctx)
		return "", -3, fmt.Errorf("ExecRepo - Save - failed to commit: %w", err)
	}

	errcode := 1
	if len(idList) > 0 {
		errcode = 0
	}

	return idList, errcode, nil
}

// Query
func (r *UserRepo) Query(ctx context.Context, table string, fieldList []string) (string, int, error) {
	userID := ""
	sql, args, err := r.Builder.Select(fieldList...).From(table).ToSql()
	if err != nil {
		return "", -1, fmt.Errorf("ExecRepo - Query - r.Builder: %w", err)
	}

	count := 0
	err = r.Pool.QueryRow(ctx, sql, args...).Scan(&count)
	if err != nil {
		return "", -2, fmt.Errorf("ExecRepo - Query - r.Pool.Query: %w", err)
	}

	errcode := 1
	if count > 0 {
		errcode = 0
	} else {
		//
	}

	return userID, errcode, nil
}
