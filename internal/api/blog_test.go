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
				requireBodyMatchAccount(t, recorder.Body, blog)
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
		// {
		// 	name:   "InvalidID",
		// 	blogID: uuid.UUID{},
		// 	buildStubs: func(store *mockdb.MockStore) {
		// 		// build stubs
		// 		store.EXPECT().
		// 			GetBlog(gomock.Any(), gomock.Any()).
		// 			Times(1).
		// 			Return(db.Blog{}, db.ErrRecordNotFound)
		// 	},
		// 	checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
		// 		// check response
		// 		require.Equal(t, http.StatusBadRequest, recorder.Code)
		// 	},
		// },
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
func requireBodyMatchAccount(t *testing.T, body *bytes.Buffer, blog db.Blog) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)

	var queriedBlog db.Blog
	err = json.Unmarshal(data, &queriedBlog)
	require.NoError(t, err)
	require.Equal(t, blog, queriedBlog)
}
