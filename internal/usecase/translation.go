package usecase

import (
	"context"
	"fmt"

	"github.com/HUSTtoKTH/lintserver/internal/entity"
	"github.com/HUSTtoKTH/lintserver/pkg/comerr"
)

// LintUseCase -.
type LintUseCase struct {
	repo   LintRepo
	webAPI AccountWebAPI
}

// New -.
func New(r LintRepo, w AccountWebAPI) *LintUseCase {
	return &LintUseCase{
		repo:   r,
		webAPI: w,
	}
}

// Upload TODO
func (uc *LintUseCase) Upload(ctx context.Context, l entity.Lint, token string) error {
	userId, err := uc.webAPI.Verify(ctx, token)
	if err != nil {
		return fmt.Errorf("LintUseCase - Upload - s.webAPI.IsAdmin: %v,%w", err, comerr.ErrUnauthorized)
	}
	// check if user is admin
	isAdmin, err := uc.webAPI.IsAdmin(ctx, userId, l.ProjectId)
	if err != nil {
		return fmt.Errorf("LintUseCase - Upload - s.webAPI.IsAdmin: %w", err)
	}
	if !isAdmin {
		// return error
		return fmt.Errorf("LintUseCase - Upload - s.webAPI.IsAdmin: %w", comerr.ErrPermission)
	} else {
		err := uc.repo.Upsert(ctx, l)
		if err != nil {
			return fmt.Errorf("LintUseCase - Upload - s.repo.Store: %w", err)
		}
		return nil
	}
}

// GetRule TODO
func (uc *LintUseCase) GetRule(ctx context.Context, projectId int64, token string) (*entity.Lint, error) {
	userId, err := uc.webAPI.Verify(ctx, token)
	if err != nil {
		return nil, fmt.Errorf("LintUseCase - GetRule - s.webAPI.Verify: %v,%w", err, comerr.ErrUnauthorized)
	}
	// check if user has project
	projects, err := uc.webAPI.GetUserProjects(ctx, userId)
	if err != nil {
		return nil, fmt.Errorf("LintUseCase - GetRule - s.webAPI.GetUserProjects: %w", err)
	}
	var hasProject bool
	for _, project := range projects {
		if project == projectId {
			hasProject = true
			break
		}
	}
	if !hasProject {
		return nil, fmt.Errorf("LintUseCase - GetRule - s.webAPI.GetUserProjects: %w", comerr.ErrPermission)
	} else {
		res, err := uc.repo.GetRule(context.Background(), projectId)
		if err != nil {
			return nil, fmt.Errorf("LintUseCase - GetRule - s.repo.GetRule: %w", err)
		}
		return res, nil
	}
}
