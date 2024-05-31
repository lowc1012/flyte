// Code generated by mockery v1.0.1. DO NOT EDIT.

package mocks

import (
	context "context"

	github "github.com/google/go-github/v42/github"
	mock "github.com/stretchr/testify/mock"
)

// GHRepoService is an autogenerated mock type for the GHRepoService type
type GHRepoService struct {
	mock.Mock
}

type GHRepoService_GetCommitSHA1 struct {
	*mock.Call
}

func (_m GHRepoService_GetCommitSHA1) Return(_a0 string, _a1 *github.Response, _a2 error) *GHRepoService_GetCommitSHA1 {
	return &GHRepoService_GetCommitSHA1{Call: _m.Call.Return(_a0, _a1, _a2)}
}

func (_m *GHRepoService) OnGetCommitSHA1(ctx context.Context, owner string, repo string, ref string, lastSHA string) *GHRepoService_GetCommitSHA1 {
	c_call := _m.On("GetCommitSHA1", ctx, owner, repo, ref, lastSHA)
	return &GHRepoService_GetCommitSHA1{Call: c_call}
}

func (_m *GHRepoService) OnGetCommitSHA1Match(matchers ...interface{}) *GHRepoService_GetCommitSHA1 {
	c_call := _m.On("GetCommitSHA1", matchers...)
	return &GHRepoService_GetCommitSHA1{Call: c_call}
}

// GetCommitSHA1 provides a mock function with given fields: ctx, owner, repo, ref, lastSHA
func (_m *GHRepoService) GetCommitSHA1(ctx context.Context, owner string, repo string, ref string, lastSHA string) (string, *github.Response, error) {
	ret := _m.Called(ctx, owner, repo, ref, lastSHA)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, string) string); ok {
		r0 = rf(ctx, owner, repo, ref, lastSHA)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 *github.Response
	if rf, ok := ret.Get(1).(func(context.Context, string, string, string, string) *github.Response); ok {
		r1 = rf(ctx, owner, repo, ref, lastSHA)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*github.Response)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, string, string, string, string) error); ok {
		r2 = rf(ctx, owner, repo, ref, lastSHA)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

type GHRepoService_GetLatestRelease struct {
	*mock.Call
}

func (_m GHRepoService_GetLatestRelease) Return(_a0 *github.RepositoryRelease, _a1 *github.Response, _a2 error) *GHRepoService_GetLatestRelease {
	return &GHRepoService_GetLatestRelease{Call: _m.Call.Return(_a0, _a1, _a2)}
}

func (_m *GHRepoService) OnGetLatestRelease(ctx context.Context, owner string, repo string) *GHRepoService_GetLatestRelease {
	c_call := _m.On("GetLatestRelease", ctx, owner, repo)
	return &GHRepoService_GetLatestRelease{Call: c_call}
}

func (_m *GHRepoService) OnGetLatestReleaseMatch(matchers ...interface{}) *GHRepoService_GetLatestRelease {
	c_call := _m.On("GetLatestRelease", matchers...)
	return &GHRepoService_GetLatestRelease{Call: c_call}
}

// GetLatestRelease provides a mock function with given fields: ctx, owner, repo
func (_m *GHRepoService) GetLatestRelease(ctx context.Context, owner string, repo string) (*github.RepositoryRelease, *github.Response, error) {
	ret := _m.Called(ctx, owner, repo)

	var r0 *github.RepositoryRelease
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *github.RepositoryRelease); ok {
		r0 = rf(ctx, owner, repo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*github.RepositoryRelease)
		}
	}

	var r1 *github.Response
	if rf, ok := ret.Get(1).(func(context.Context, string, string) *github.Response); ok {
		r1 = rf(ctx, owner, repo)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*github.Response)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, string, string) error); ok {
		r2 = rf(ctx, owner, repo)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

type GHRepoService_GetReleaseByTag struct {
	*mock.Call
}

func (_m GHRepoService_GetReleaseByTag) Return(_a0 *github.RepositoryRelease, _a1 *github.Response, _a2 error) *GHRepoService_GetReleaseByTag {
	return &GHRepoService_GetReleaseByTag{Call: _m.Call.Return(_a0, _a1, _a2)}
}

func (_m *GHRepoService) OnGetReleaseByTag(ctx context.Context, owner string, repo string, tag string) *GHRepoService_GetReleaseByTag {
	c_call := _m.On("GetReleaseByTag", ctx, owner, repo, tag)
	return &GHRepoService_GetReleaseByTag{Call: c_call}
}

func (_m *GHRepoService) OnGetReleaseByTagMatch(matchers ...interface{}) *GHRepoService_GetReleaseByTag {
	c_call := _m.On("GetReleaseByTag", matchers...)
	return &GHRepoService_GetReleaseByTag{Call: c_call}
}

// GetReleaseByTag provides a mock function with given fields: ctx, owner, repo, tag
func (_m *GHRepoService) GetReleaseByTag(ctx context.Context, owner string, repo string, tag string) (*github.RepositoryRelease, *github.Response, error) {
	ret := _m.Called(ctx, owner, repo, tag)

	var r0 *github.RepositoryRelease
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) *github.RepositoryRelease); ok {
		r0 = rf(ctx, owner, repo, tag)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*github.RepositoryRelease)
		}
	}

	var r1 *github.Response
	if rf, ok := ret.Get(1).(func(context.Context, string, string, string) *github.Response); ok {
		r1 = rf(ctx, owner, repo, tag)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*github.Response)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, string, string, string) error); ok {
		r2 = rf(ctx, owner, repo, tag)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

type GHRepoService_ListReleases struct {
	*mock.Call
}

func (_m GHRepoService_ListReleases) Return(_a0 []*github.RepositoryRelease, _a1 *github.Response, _a2 error) *GHRepoService_ListReleases {
	return &GHRepoService_ListReleases{Call: _m.Call.Return(_a0, _a1, _a2)}
}

func (_m *GHRepoService) OnListReleases(ctx context.Context, owner string, repo string, opts *github.ListOptions) *GHRepoService_ListReleases {
	c_call := _m.On("ListReleases", ctx, owner, repo, opts)
	return &GHRepoService_ListReleases{Call: c_call}
}

func (_m *GHRepoService) OnListReleasesMatch(matchers ...interface{}) *GHRepoService_ListReleases {
	c_call := _m.On("ListReleases", matchers...)
	return &GHRepoService_ListReleases{Call: c_call}
}

// ListReleases provides a mock function with given fields: ctx, owner, repo, opts
func (_m *GHRepoService) ListReleases(ctx context.Context, owner string, repo string, opts *github.ListOptions) ([]*github.RepositoryRelease, *github.Response, error) {
	ret := _m.Called(ctx, owner, repo, opts)

	var r0 []*github.RepositoryRelease
	if rf, ok := ret.Get(0).(func(context.Context, string, string, *github.ListOptions) []*github.RepositoryRelease); ok {
		r0 = rf(ctx, owner, repo, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*github.RepositoryRelease)
		}
	}

	var r1 *github.Response
	if rf, ok := ret.Get(1).(func(context.Context, string, string, *github.ListOptions) *github.Response); ok {
		r1 = rf(ctx, owner, repo, opts)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*github.Response)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, string, string, *github.ListOptions) error); ok {
		r2 = rf(ctx, owner, repo, opts)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}
