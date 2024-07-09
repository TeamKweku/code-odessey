// Package api implements the client-side API for code wishing to interact
// with the code-odessey service. The methods of the [Server] type correspond
// to the code-odessey REST API.
//
// # Examples
//
// Several examples of using this package are available [in the GitHub
// repository].
//
// [in the GitHub repository]: https://github.com/teamkweku/code-odessey/tree/main/examples
//
// [the API documentation]: https://github.com/teamkweku/code-odessey/blob/main/docs/api.md
package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/teamkweku/code-odessey/config"
	db "github.com/teamkweku/code-odessey/internal/db/sqlc"
	"github.com/teamkweku/code-odessey/internal/token"
)

// Server encapsulates the db, Engine, config
// to help connect to db and its related functions
// Engine helps send api request to correct handler for
// processing
type Server struct {
	store      db.Store
	router     *gin.Engine
	tokenMaker token.Maker
	config     config.Config
}

// create a new server instance
func NewServer(config config.Config, store db.Store) (*Server, error) {
	// either create a PASETO OR JWT token maker
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	server.setupRouter()

	return server, nil
}

// setupRouter sets up the routing for the server.
func (server *Server) setupRouter() {
	router := gin.Default()

	// user routes
	router.POST("/users", server.createUser)
	router.POST("users/login", server.loginUser)
	router.POST("/tokens/renew_access", server.renewAccessToken)

	router.GET("/blogs", server.listBlogs)
	router.GET("/blogs/:id", server.getBlogByID)

	// adding a routes group
	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))

	authRoutes.POST("/blogs", server.createBlog)

	authRoutes.PUT("/blogs/:id", server.updateBlog)
	authRoutes.DELETE("/blogs/:id", server.deleteBlog)

	// Listing comments for a specific blog
	authRoutes.GET("/blogs/:id/comments", server.listCommentsbyBlogID)

	// update comment of a specific blog id
	authRoutes.PUT("/blogs/:id/comments/:comment_id", server.updateCommentByBlogID)

	authRoutes.DELETE("/blogs/:id/comments/:comment_id", server.deleteCommentByBlogID)

	// creating comments
	authRoutes.POST("/comments", server.createComment)
	authRoutes.GET("/comments/:id", server.getCommentByID)

	server.router = router
}

// Start runs the HTTP server on a specific address.
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

// errorResponse formats an error message as a JSON response.
func errorResponse(err error) gin.H {
	// return gin.H{"error": err.Error()}
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		errMsgs := make([]string, len(validationErrors))
		for i, fieldError := range validationErrors {
			errMsgs[i] = fmt.Sprintf("Error: Field validation for '%s' failed on the '%s' tag", fieldError.Field(), fieldError.Tag())
		}
		return gin.H{"errors": errMsgs}
	}
	return gin.H{"error": err.Error()}
}
