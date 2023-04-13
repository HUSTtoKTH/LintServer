package repo

import (
	"context"
	"fmt"

	"github.com/HUSTtoKTH/lintserver/internal/entity"
	"github.com/HUSTtoKTH/lintserver/pkg/comerr"
	"github.com/HUSTtoKTH/lintserver/pkg/postgres"
)

// LintRepo TODO
type LintRepo struct {
	*postgres.Postgres
}

// New -.
func New(pg *postgres.Postgres) *LintRepo {
	return &LintRepo{pg}
}

// Upsert TODO
// Store -.
func (r *LintRepo) Upsert(ctx context.Context, t entity.Lint) error {
	sql, args, err := r.Builder.
		Insert("rules").
		Columns("project_id, organization_id, rule").
		Values(t.ProjectId, t.OrganizationId, t.Rule).
		Suffix("ON CONFLICT (project_id) DO UPDATE SET rule = ?", t.Rule).
		ToSql()
	if err != nil {
		return fmt.Errorf("LintRepo - Store - r.Builder: %w", err)
	}

	_, err = r.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("LintRepo - Store - r.Pool.Exec: %w", err)
	}

	return nil
}

// GetRule TODO
func (r *LintRepo) GetRule(ctx context.Context, projectId int64) (*entity.Lint, error) {
	sql, args, err := r.Builder.
		Select("project_id, organization_id, rule").
		From("rules").
		Where("project_id = ?", projectId).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("LintRepo - GetRule - r.Builder: %w", err)
	}

	rows, err := r.Pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, fmt.Errorf("LintRepo - GetRule - r.Pool.Query: %w", err)
	}
	defer rows.Close()

	entities := []*entity.Lint{}

	for rows.Next() {
		e := &entity.Lint{}
		err = rows.Scan(&e.OrganizationId, &e.ProjectId, &e.Rule)
		if err != nil {
			return nil, fmt.Errorf("LintRepo - GetRule - rows.Scan: %w", err)
		}
		entities = append(entities, e)
	}
	if len(entities) == 0 {
		return nil, fmt.Errorf("LintRepo - GetRule - %w", comerr.ErrNoRecord)
	}
	return entities[0], nil
}
