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
)

// Server encapsulates the db, Engine, config
// to help connect to db and its related functions
// Engine helps send api request to correct handler for
// processing
type Server struct {
	store  db.Store
	router *gin.Engine
	config config.Config
}

// create a new server instance
func NewServer(config config.Config, store db.Store) (*Server, error) {
	server := &Server{
		config: config,
		store:  store,
	}

	server.setupRouter()

	return server, nil
}

// setupRouter sets up the routing for the server.
func (server *Server) setupRouter() {
	router := gin.Default()

	// creating the post request to create a blog
	router.POST("/blogs", server.createBlog)
	router.GET("/blogs/:id", server.getBlogByID)
	router.GET("/blogs", server.listBlogs)
	router.PUT("/blogs/:id", server.updateBlog)
	router.DELETE("/blogs/:id", server.deleteBlog)

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
