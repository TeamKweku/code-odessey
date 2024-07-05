package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	mockdb "github.com/teamkweku/code-odessey/internal/db/mock"
	db "github.com/teamkweku/code-odessey/internal/db/sqlc"
	"github.com/teamkweku/code-odessey/pkg/utils"
	"go.uber.org/mock/gomock"
)

func TestGetBlogAPI(t *testing.T) {
	blog := randomBlog()

	// testing for different cases
	testCases := []struct {
		name          string
		blogID        uuid.UUID
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name:   "StatusOK",
			blogID: blog.ID,
			buildStubs: func(store *mockdb.MockStore) {
				// build stubs
				store.EXPECT().
					GetBlog(gomock.Any(), gomock.Eq(blog.ID)).
					Times(1).
					Return(blog, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				// check response
				require.Equal(t, http.StatusOK, recorder.Code)

				// check body of response
				requireBodyMatchBlog(t, recorder.Body, blog)
			},
		},
		{
			name:   "NotFound",
			blogID: blog.ID,
			buildStubs: func(store *mockdb.MockStore) {
				// build stubs
				store.EXPECT().
					GetBlog(gomock.Any(), gomock.Eq(blog.ID)).
					Times(1).
					Return(db.Blog{}, db.ErrRecordNotFound)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				// check response
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name:   "InternalServerError",
			blogID: blog.ID,
			buildStubs: func(store *mockdb.MockStore) {
				// build stubs
				store.EXPECT().
					GetBlog(gomock.Any(), gomock.Eq(blog.ID)).
					Times(1).
					Return(db.Blog{}, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				// check response
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name:   "InvalidID",
			blogID: uuid.UUID{},
			buildStubs: func(store *mockdb.MockStore) {
				// build stubs
				store.EXPECT().
					GetBlog(gomock.Any(), gomock.Any()).
					Times(0)

			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				// check response
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
	}

	// testing all cases in the testCases
	// running each case as a seperate test with the Run() func
	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			// Creating an equivalent store object to the store struct
			// to be able to mock db and api calls
			store := mockdb.NewMockStore(ctrl)
			tc.buildStubs(store)

			// start a test server and send request
			server := newTestServer(t, store)
			recorder := httptest.NewRecorder()

			url := fmt.Sprintf("/blogs/%s", tc.blogID)
			// making a test request to the getBlog url
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})
	}
}

// Testing CreateBlog API route
func TestCreateBlogAPI(t *testing.T) {
	blog := randomBlog()

	testCases := []struct {
		name          string
		body          gin.H
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			body: gin.H{
				"title":        blog.Title,
				"slug":         blog.Slug,
				"description":  blog.Description,
				"body":         blog.Body,
				"banner_image": blog.BannerImage,
			},
			buildStubs: func(store *mockdb.MockStore) {
				arg := db.CreateBlogParams{
					Title:       blog.Title,
					Slug:        blog.Slug,
					Description: blog.Description,
					Body:        blog.Body,
					BannerImage: blog.BannerImage,
				}

				store.EXPECT().
					CreateBlog(gomock.Any(), gomock.Eq(arg)).
					Times(1).
					Return(blog, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusCreated, recorder.Code)
				requireBodyMatchBlog(t, recorder.Body, blog)
			},
		},
		{
			name: "InvalidRequest",
			body: gin.H{},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					CreateBlog(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "InternalError",
			body: gin.H{
				"title":        blog.Title,
				"slug":         blog.Slug,
				"description":  blog.Description,
				"body":         blog.Body,
				"banner_image": blog.BannerImage,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					CreateBlog(gomock.Any(), gomock.Any()).
					Times(1).
					Return(db.Blog{}, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)
			tc.buildStubs(store)

			server := newTestServer(t, store)
			recorder := httptest.NewRecorder()

			data, err := json.Marshal(tc.body)
			require.NoError(t, err)

			url := "/blogs"
			request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})
	}
}

// Tesing the ListBlogAPi
func TestListBlogsAPI(t *testing.T) {
	n := 5
	blogs := make([]db.Blog, n)
	for i := 0; i < n; i++ {
		blogs[i] = randomBlog()
	}

	type Query struct {
		pageID   int
		pageSize int
	}

	testCases := []struct {
		name          string
		query         Query
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			query: Query{
				pageID:   1,
				pageSize: n,
			},

			buildStubs: func(store *mockdb.MockStore) {
				arg := db.ListBlogsParams{
					Limit:  5,
					Offset: 0,
				}

				store.EXPECT().
					ListBlogs(gomock.Any(), gomock.Eq(arg)).
					Times(1).
					Return(blogs, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchBlogs(t, recorder.Body, blogs)
			},
		},
		{
			name: "InvalidPageID",
			query: Query{
				pageID:   -1,
				pageSize: n,
			},

			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					ListBlogs(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "InvalidPageSize",
			query: Query{
				pageID:   1,
				pageSize: 200,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					ListBlogs(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "InternalError",
			query: Query{
				pageID:   1,
				pageSize: n,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					ListBlogs(gomock.Any(), gomock.Any()).
					Times(1).
					Return([]db.Blog{}, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)
			tc.buildStubs(store)

			server := newTestServer(t, store)
			recorder := httptest.NewRecorder()

			url := "/blogs"
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)

			// Add query parameters to request URL
			q := request.URL.Query()
			q.Add("page_id", fmt.Sprintf("%d", tc.query.pageID))
			q.Add("page_size", fmt.Sprintf("%d", tc.query.pageSize))
			request.URL.RawQuery = q.Encode()

			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})
	}
}

// Generate random blog to mockdb
func randomBlog() db.Blog {
	return db.Blog{
		ID:          uuid.New(),
		Title:       utils.RandomTitle(),
		Slug:        utils.RandomSlug() + "-" + uuid.New().String(),
		Description: utils.RandomDescription(),
		BannerImage: utils.RandomImageURL(),
		Body:        utils.RandomParagraph(),
	}
}

// function to compare mock Response Body and created blog body
func requireBodyMatchBlog(t *testing.T, body *bytes.Buffer, blog db.Blog) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)

	var queriedBlog db.Blog
	err = json.Unmarshal(data, &queriedBlog)
	require.NoError(t, err)
	require.Equal(t, blog, queriedBlog)
}

// function to compore the mock Reponse with the list of blogs fetched
func requireBodyMatchBlogs(t *testing.T, body *bytes.Buffer, blogs []db.Blog) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)

	var queriedBlog []db.Blog
	err = json.Unmarshal(data, &queriedBlog)
	require.NoError(t, err)
	require.Equal(t, blogs, queriedBlog)
}
