// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/teamkweku/code-odessey/internal/db/sqlc (interfaces: Store)
//
// Generated by this command:
//
//	mockgen -package mockdb -destination internal/db/mock/store.go github.com/teamkweku/code-odessey/internal/db/sqlc Store
//

// Package mockdb is a generated GoMock package.
package mockdb

import (
	context "context"
	reflect "reflect"

	uuid "github.com/google/uuid"
	pgconn "github.com/jackc/pgx/v5/pgconn"
	db "github.com/teamkweku/code-odessey/internal/db/sqlc"
	gomock "go.uber.org/mock/gomock"
)

// MockStore is a mock of Store interface.
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore.
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance.
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// CreateBlog mocks base method.
func (m *MockStore) CreateBlog(arg0 context.Context, arg1 db.CreateBlogParams) (db.Blog, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateBlog", arg0, arg1)
	ret0, _ := ret[0].(db.Blog)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateBlog indicates an expected call of CreateBlog.
func (mr *MockStoreMockRecorder) CreateBlog(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBlog", reflect.TypeOf((*MockStore)(nil).CreateBlog), arg0, arg1)
}

// CreateComment mocks base method.
func (m *MockStore) CreateComment(arg0 context.Context, arg1 db.CreateCommentParams) (db.Comment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateComment", arg0, arg1)
	ret0, _ := ret[0].(db.Comment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateComment indicates an expected call of CreateComment.
func (mr *MockStoreMockRecorder) CreateComment(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateComment", reflect.TypeOf((*MockStore)(nil).CreateComment), arg0, arg1)
}

// CreateFavorite mocks base method.
func (m *MockStore) CreateFavorite(arg0 context.Context, arg1 db.CreateFavoriteParams) (db.Favorite, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateFavorite", arg0, arg1)
	ret0, _ := ret[0].(db.Favorite)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateFavorite indicates an expected call of CreateFavorite.
func (mr *MockStoreMockRecorder) CreateFavorite(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateFavorite", reflect.TypeOf((*MockStore)(nil).CreateFavorite), arg0, arg1)
}

// CreateSession mocks base method.
func (m *MockStore) CreateSession(arg0 context.Context, arg1 db.CreateSessionParams) (db.Session, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSession", arg0, arg1)
	ret0, _ := ret[0].(db.Session)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSession indicates an expected call of CreateSession.
func (mr *MockStoreMockRecorder) CreateSession(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSession", reflect.TypeOf((*MockStore)(nil).CreateSession), arg0, arg1)
}

// CreateUser mocks base method.
func (m *MockStore) CreateUser(arg0 context.Context, arg1 db.CreateUserParams) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockStoreMockRecorder) CreateUser(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockStore)(nil).CreateUser), arg0, arg1)
}

// DeleteBlog mocks base method.
func (m *MockStore) DeleteBlog(arg0 context.Context, arg1 uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteBlog", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteBlog indicates an expected call of DeleteBlog.
func (mr *MockStoreMockRecorder) DeleteBlog(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteBlog", reflect.TypeOf((*MockStore)(nil).DeleteBlog), arg0, arg1)
}

// DeleteBlogTx mocks base method.
func (m *MockStore) DeleteBlogTx(arg0 context.Context, arg1 db.DeleteBlogTxParams) (db.DeleteBlogTxResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteBlogTx", arg0, arg1)
	ret0, _ := ret[0].(db.DeleteBlogTxResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteBlogTx indicates an expected call of DeleteBlogTx.
func (mr *MockStoreMockRecorder) DeleteBlogTx(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteBlogTx", reflect.TypeOf((*MockStore)(nil).DeleteBlogTx), arg0, arg1)
}

// DeleteComment mocks base method.
func (m *MockStore) DeleteComment(arg0 context.Context, arg1 uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteComment", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteComment indicates an expected call of DeleteComment.
func (mr *MockStoreMockRecorder) DeleteComment(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteComment", reflect.TypeOf((*MockStore)(nil).DeleteComment), arg0, arg1)
}

// DeleteCommentByBlogID mocks base method.
func (m *MockStore) DeleteCommentByBlogID(arg0 context.Context, arg1 db.DeleteCommentByBlogIDParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCommentByBlogID", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCommentByBlogID indicates an expected call of DeleteCommentByBlogID.
func (mr *MockStoreMockRecorder) DeleteCommentByBlogID(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCommentByBlogID", reflect.TypeOf((*MockStore)(nil).DeleteCommentByBlogID), arg0, arg1)
}

// DeleteCommentsByBlog mocks base method.
func (m *MockStore) DeleteCommentsByBlog(arg0 context.Context, arg1 uuid.UUID) (pgconn.CommandTag, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCommentsByBlog", arg0, arg1)
	ret0, _ := ret[0].(pgconn.CommandTag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteCommentsByBlog indicates an expected call of DeleteCommentsByBlog.
func (mr *MockStoreMockRecorder) DeleteCommentsByBlog(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCommentsByBlog", reflect.TypeOf((*MockStore)(nil).DeleteCommentsByBlog), arg0, arg1)
}

// DeleteFavorite mocks base method.
func (m *MockStore) DeleteFavorite(arg0 context.Context, arg1 uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteFavorite", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteFavorite indicates an expected call of DeleteFavorite.
func (mr *MockStoreMockRecorder) DeleteFavorite(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteFavorite", reflect.TypeOf((*MockStore)(nil).DeleteFavorite), arg0, arg1)
}

// DeleteFavoritesByBlog mocks base method.
func (m *MockStore) DeleteFavoritesByBlog(arg0 context.Context, arg1 uuid.UUID) (pgconn.CommandTag, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteFavoritesByBlog", arg0, arg1)
	ret0, _ := ret[0].(pgconn.CommandTag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteFavoritesByBlog indicates an expected call of DeleteFavoritesByBlog.
func (mr *MockStoreMockRecorder) DeleteFavoritesByBlog(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteFavoritesByBlog", reflect.TypeOf((*MockStore)(nil).DeleteFavoritesByBlog), arg0, arg1)
}

// GetBlog mocks base method.
func (m *MockStore) GetBlog(arg0 context.Context, arg1 uuid.UUID) (db.Blog, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBlog", arg0, arg1)
	ret0, _ := ret[0].(db.Blog)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBlog indicates an expected call of GetBlog.
func (mr *MockStoreMockRecorder) GetBlog(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlog", reflect.TypeOf((*MockStore)(nil).GetBlog), arg0, arg1)
}

// GetBlogBySlug mocks base method.
func (m *MockStore) GetBlogBySlug(arg0 context.Context, arg1 string) (db.Blog, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBlogBySlug", arg0, arg1)
	ret0, _ := ret[0].(db.Blog)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBlogBySlug indicates an expected call of GetBlogBySlug.
func (mr *MockStoreMockRecorder) GetBlogBySlug(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlogBySlug", reflect.TypeOf((*MockStore)(nil).GetBlogBySlug), arg0, arg1)
}

// GetComment mocks base method.
func (m *MockStore) GetComment(arg0 context.Context, arg1 uuid.UUID) (db.Comment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetComment", arg0, arg1)
	ret0, _ := ret[0].(db.Comment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetComment indicates an expected call of GetComment.
func (mr *MockStoreMockRecorder) GetComment(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetComment", reflect.TypeOf((*MockStore)(nil).GetComment), arg0, arg1)
}

// GetFavorite mocks base method.
func (m *MockStore) GetFavorite(arg0 context.Context, arg1 uuid.UUID) (db.Favorite, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFavorite", arg0, arg1)
	ret0, _ := ret[0].(db.Favorite)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFavorite indicates an expected call of GetFavorite.
func (mr *MockStoreMockRecorder) GetFavorite(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFavorite", reflect.TypeOf((*MockStore)(nil).GetFavorite), arg0, arg1)
}

// GetSession mocks base method.
func (m *MockStore) GetSession(arg0 context.Context, arg1 uuid.UUID) (db.Session, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSession", arg0, arg1)
	ret0, _ := ret[0].(db.Session)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSession indicates an expected call of GetSession.
func (mr *MockStoreMockRecorder) GetSession(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSession", reflect.TypeOf((*MockStore)(nil).GetSession), arg0, arg1)
}

// GetUser mocks base method.
func (m *MockStore) GetUser(arg0 context.Context, arg1 uuid.UUID) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockStoreMockRecorder) GetUser(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockStore)(nil).GetUser), arg0, arg1)
}

// GetUserByUsername mocks base method.
func (m *MockStore) GetUserByUsername(arg0 context.Context, arg1 string) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByUsername", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByUsername indicates an expected call of GetUserByUsername.
func (mr *MockStoreMockRecorder) GetUserByUsername(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByUsername", reflect.TypeOf((*MockStore)(nil).GetUserByUsername), arg0, arg1)
}

// ListBlogs mocks base method.
func (m *MockStore) ListBlogs(arg0 context.Context, arg1 db.ListBlogsParams) ([]db.Blog, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListBlogs", arg0, arg1)
	ret0, _ := ret[0].([]db.Blog)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListBlogs indicates an expected call of ListBlogs.
func (mr *MockStoreMockRecorder) ListBlogs(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListBlogs", reflect.TypeOf((*MockStore)(nil).ListBlogs), arg0, arg1)
}

// ListCommentsByBlog mocks base method.
func (m *MockStore) ListCommentsByBlog(arg0 context.Context, arg1 db.ListCommentsByBlogParams) ([]db.Comment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListCommentsByBlog", arg0, arg1)
	ret0, _ := ret[0].([]db.Comment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListCommentsByBlog indicates an expected call of ListCommentsByBlog.
func (mr *MockStoreMockRecorder) ListCommentsByBlog(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListCommentsByBlog", reflect.TypeOf((*MockStore)(nil).ListCommentsByBlog), arg0, arg1)
}

// ListFavoritesByBlog mocks base method.
func (m *MockStore) ListFavoritesByBlog(arg0 context.Context, arg1 db.ListFavoritesByBlogParams) ([]db.Favorite, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListFavoritesByBlog", arg0, arg1)
	ret0, _ := ret[0].([]db.Favorite)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListFavoritesByBlog indicates an expected call of ListFavoritesByBlog.
func (mr *MockStoreMockRecorder) ListFavoritesByBlog(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListFavoritesByBlog", reflect.TypeOf((*MockStore)(nil).ListFavoritesByBlog), arg0, arg1)
}

// UpdateBlog mocks base method.
func (m *MockStore) UpdateBlog(arg0 context.Context, arg1 db.UpdateBlogParams) (db.Blog, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateBlog", arg0, arg1)
	ret0, _ := ret[0].(db.Blog)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateBlog indicates an expected call of UpdateBlog.
func (mr *MockStoreMockRecorder) UpdateBlog(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateBlog", reflect.TypeOf((*MockStore)(nil).UpdateBlog), arg0, arg1)
}

// UpdateComment mocks base method.
func (m *MockStore) UpdateComment(arg0 context.Context, arg1 db.UpdateCommentParams) (db.Comment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateComment", arg0, arg1)
	ret0, _ := ret[0].(db.Comment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateComment indicates an expected call of UpdateComment.
func (mr *MockStoreMockRecorder) UpdateComment(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateComment", reflect.TypeOf((*MockStore)(nil).UpdateComment), arg0, arg1)
}

// UpdateCommentByBlogID mocks base method.
func (m *MockStore) UpdateCommentByBlogID(arg0 context.Context, arg1 db.UpdateCommentByBlogIDParams) (db.Comment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCommentByBlogID", arg0, arg1)
	ret0, _ := ret[0].(db.Comment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateCommentByBlogID indicates an expected call of UpdateCommentByBlogID.
func (mr *MockStoreMockRecorder) UpdateCommentByBlogID(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCommentByBlogID", reflect.TypeOf((*MockStore)(nil).UpdateCommentByBlogID), arg0, arg1)
}

// UpdateUser mocks base method.
func (m *MockStore) UpdateUser(arg0 context.Context, arg1 db.UpdateUserParams) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockStoreMockRecorder) UpdateUser(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockStore)(nil).UpdateUser), arg0, arg1)
}
