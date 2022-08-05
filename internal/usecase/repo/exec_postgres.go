package repo

import (
	"context"
	"fmt"

	"gexec/internal/entity"
	"gexec/pkg/postgres"
)

const _defaultEntityCap = 64

// ValueRepo -.
type ValueRepo struct {
	*postgres.Postgres
}

// New -.
func New(pg *postgres.Postgres) *ValueRepo {
	return &ValueRepo{pg}
}

// Save -.
func (r *ValueRepo) Save(ctx context.Context, entities []entity.TableEntity) (map[string][]string, int, error) {
	idMap := make(map[string][]string)
	tx, err := r.Pool.Begin(ctx)
	if err != nil {
		return idMap, -3, fmt.Errorf("ExecRepo - Save - r.Pool.Begin: %w", err)
	}
	for _, entity := range entities {
		builder := r.Builder.Insert(entity.Table).Columns(entity.Columns...)
		for _, value := range entity.Values {
			builder = builder.Values(value...)
		}
		sql, args, err := builder.Suffix("RETURNING id").ToSql()
		if err != nil {
			return idMap, -1, fmt.Errorf("ExecRepo - Save - r.Builder: %w", err)
		}

		var idList []string
		err = tx.QueryRow(ctx, sql, args...).Scan(&idList)
		if err != nil {
			tx.Rollback(ctx)
			return idMap, -2, fmt.Errorf("ExecRepo - Save r.Pool.Query: %w", err)
		}
		idMap[entity.Table] = idList
	}

	err = tx.Commit(ctx)
	if err != nil {
		tx.Rollback(ctx)
		return idMap, -3, fmt.Errorf("ExecRepo - Save - failed to commit: %w", err)
	}

	return idMap, 0, nil
}

// Query
func (r *ValueRepo) Query(ctx context.Context, table string, fieldList []string) (string, int, error) {
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
