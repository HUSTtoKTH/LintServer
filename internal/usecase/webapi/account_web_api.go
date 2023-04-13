package webapi

import (
	"context"
	"errors"
)

// AccountWebAPI TODO
// LintWebAPI -.
type AccountWebAPI struct {
}

// New -.
func New() *AccountWebAPI {
	return &AccountWebAPI{}
}

// Verify mock api
func (t *AccountWebAPI) Verify(ctx context.Context, token string) (userId int64, err error) {
	if token == "superadmin" {
		return 1, nil
	} else if token != "" {
		return 2, nil
	} else {
		return 0, errors.New("invalid token")
	}
}

// IsAdmin mock api
func (t *AccountWebAPI) IsAdmin(ctx context.Context, userId int64, organizationId int64) (bool, error) {
	if userId == 1 {
		return true, nil
	}
	return false, nil
}

// GetUserProjects mock api
func (t *AccountWebAPI) GetUserProjects(ctx context.Context, userId int64) ([]int64, error) {
	if userId == 1 {
		return []int64{1, 2}, nil
	} else {
		return []int64{1}, nil
	}
}
