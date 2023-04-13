package usecase_test

import (
	"context"
	"errors"
	"testing"

	"github.com/HUSTtoKTH/lintserver/internal/entity"
	"github.com/HUSTtoKTH/lintserver/internal/usecase"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

var errInternalServErr = errors.New("internal server error")

type test struct {
	name string
	mock func()
	res  interface{}
	err  error
}

func newLint(t *testing.T) (*usecase.LintUseCase, *MockLintRepo, *MockAccountWebAPI) {
	t.Helper()

	mockCtl := gomock.NewController(t)
	defer mockCtl.Finish()

	repo := NewMockLintRepo(mockCtl)
	webAPI := NewMockAccountWebAPI(mockCtl)

	uc := usecase.New(repo, webAPI)

	return uc, repo, webAPI
}

func TestUpload(t *testing.T) {
	t.Parallel()

	uc, repo, webAPI := newLint(t)

	tests := []test{
		{
			name: "result with error",
			mock: func() {
				webAPI.EXPECT().Verify(context.Background(), "token").Return(int64(1), nil)
				webAPI.EXPECT().IsAdmin(context.Background(), int64(1), int64(1)).Return(true, nil)
				repo.EXPECT().Upsert(context.Background(), entity.Lint{
					ProjectId:      1,
					OrganizationId: 1,
					Rule:           "rule",
				}).Return(errInternalServErr)
			},
			res: nil,
			err: errInternalServErr,
		},
	}

	for _, tc := range tests {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			tc.mock()

			err := uc.Upload(context.Background(), entity.Lint{
				ProjectId:      1,
				OrganizationId: 1,
				Rule:           "rule",
			}, "token")

			require.ErrorIs(t, err, tc.err)
		})
	}
}

func TestGetRule(t *testing.T) {
	t.Parallel()

	uc, repo, webAPI := newLint(t)

	tests := []test{
		{
			name: "success result",
			mock: func() {
				webAPI.EXPECT().Verify(context.Background(), "token").Return(int64(1), nil)
				webAPI.EXPECT().GetUserProjects(context.Background(), int64(1)).Return([]int64{1, 2}, nil)
				repo.EXPECT().GetRule(context.Background(), int64(1)).Return(&entity.Lint{}, nil)
			},
			res: &entity.Lint{},
			err: nil,
		},
		{
			name: "web API error",
			mock: func() {
				webAPI.EXPECT().Verify(context.Background(), "token").Return(int64(1), nil)
				webAPI.EXPECT().GetUserProjects(context.Background(), int64(1)).Return(nil, errInternalServErr)
			},
			res: (*entity.Lint)(nil),
			err: errInternalServErr,
		},
		{
			name: "repo error",
			mock: func() {
				webAPI.EXPECT().Verify(context.Background(), "token").Return(int64(1), nil)
				webAPI.EXPECT().GetUserProjects(context.Background(), int64(1)).Return([]int64{1, 2}, nil)
				repo.EXPECT().GetRule(context.Background(), int64(1)).Return(nil, errInternalServErr)
			},
			res: (*entity.Lint)(nil),
			err: errInternalServErr,
		},
	}

	for _, tc := range tests {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			tc.mock()

			res, err := uc.GetRule(context.Background(), 1, "token")

			require.EqualValues(t, res, tc.res)
			require.ErrorIs(t, err, tc.err)
		})
	}
}
