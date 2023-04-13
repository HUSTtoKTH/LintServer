// Package usecase implements application business logic. Each logic group in own file.
package usecase

import (
	"context"

	"github.com/HUSTtoKTH/lintserver/internal/entity"
)

//go:generate mockgen -source=interfaces.go -destination=./mocks_test.go -package=usecase_test

type (
	// Lint -.
	Lint interface {
		// Translate(context.Context, entity.Lint) (entity.Lint, error)
		// History(context.Context) ([]entity.Lint, error)
		Upload(ctx context.Context, l entity.Lint, token string) error
		GetRule(ctx context.Context, projectId int64, token string) (*entity.Lint, error)
	}

	// LintRepo -.
	LintRepo interface {
		Upsert(context.Context, entity.Lint) error
		GetRule(ctx context.Context, projectId int64) (*entity.Lint, error)
	}

	// AccountWebAPI TODO
	// LintWebAPI -.
	AccountWebAPI interface {
		Verify(ctx context.Context, token string) (userId int64, err error)
		IsAdmin(ctx context.Context, userId int64, organizationId int64) (bool, error)
		GetUserProjects(ctx context.Context, userId int64) ([]int64, error)
	}
)
