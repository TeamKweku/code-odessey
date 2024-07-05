package api

import (
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	db "github.com/teamkweku/code-odessey/internal/db/sqlc"
)

// createCommentRequest represents the request structure for creating a comment
type createCommentRequest struct {
	BlogID string `json:"blog_id" binding:"required"`
	Body   string `json:"body" binding:"required,min=1"`
}

// createComment handles the creation of a new comment
func (server *Server) createComment(ctx *gin.Context) {
	var req createCommentRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		if err == io.EOF {
			ctx.JSON(http.StatusBadRequest, errorResponse(errors.New("request body is empty")))
		} else {
			ctx.JSON(http.StatusBadRequest, errorResponse(err))
		}
		return
	}

	blogID, err := uuid.Parse(req.BlogID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(fmt.Errorf("invalid blog id format")))
		return
	}

	arg := db.CreateCommentParams{
		BlogID: blogID,
		Body:   req.Body,
	}

	comment, err := server.store.CreateComment(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusCreated, comment)
}

// getCommentRequest represents the request structs for geting a comment
type getCommentRequest struct {
	ID string `uri:"id" binding:"required"`
}

// getComment handles the retrieval of a comment by its ID
func (server *Server) getCommentByID(ctx *gin.Context) {
	var req getCommentRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		if err == io.EOF {
			ctx.JSON(http.StatusBadRequest, errorResponse(errors.New("request body is empty")))
		} else {
			ctx.JSON(http.StatusBadRequest, errorResponse(err))
		}
		return
	}

	id, err := uuid.Parse(req.ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(fmt.Errorf("invalid uuid format")))
		return
	}

	comment, err := server.store.GetComment(ctx, id)
	if err != nil {
		if errors.Is(err, db.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, comment)
}

// listCommentsRequest represent the request structure for listing comments
type listCommentsRequest struct {
	PageID   int32 `form:"page_id" binding:"omitempty,min=1"`
	PageSize int32 `form:"page_limit" binding:"omitempty,min=5,max=10"`
}

// listComments handles the retrieval of a list of comments for a specific blog
func (server *Server) listCommentsbyBlogID(ctx *gin.Context) {
	var req listCommentsRequest

	if err := ctx.ShouldBindQuery(&req); err != nil {
		if err == io.EOF {
			ctx.JSON(http.StatusBadRequest, errorResponse(errors.New("request body is empty")))
		} else {
			ctx.JSON(http.StatusBadRequest, errorResponse(err))
		}
		return
	}

	ID := ctx.Param("id")
	blogID, err := uuid.Parse(ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(fmt.Errorf("invalid blog id format")))
		return
	}

	// default page_limit and page_size values when omitted
	if req.PageID == 0 {
		req.PageID = 1
	}
	if req.PageSize == 0 {
		req.PageSize = 10
	}

	arg := db.ListCommentsByBlogParams{
		BlogID: blogID,
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	comments, err := server.store.ListCommentsByBlog(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, comments)
}

// updateBlogIDCommentRequest handles the update of comment for a specific blog
type updateCommentByBlogIDRequest struct {
	Body string `json:"body" binding:"required"`
}

func (server *Server) updateCommentByBlogID(ctx *gin.Context) {
	var req updateCommentByBlogIDRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		if err == io.EOF {
			ctx.JSON(http.StatusBadRequest, errorResponse(errors.New("request body is empty")))
		} else {
			ctx.JSON(http.StatusBadRequest, errorResponse(err))
		}
		return
	}

	ID := ctx.Param("id")
	blogID, err := uuid.Parse(ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(fmt.Errorf("invalid blog id format")))
		return
	}

	commentID := ctx.Param("comment_id")
	commentUUID, err := uuid.Parse(commentID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(fmt.Errorf("invalid comment id format")))
		return
	}

	arg := db.UpdateCommentByBlogIDParams{
		ID:     commentUUID,
		BlogID: blogID,
		Body:   req.Body,
	}

	updatedComment, err := server.store.UpdateCommentByBlogID(ctx, arg)
	if err != nil {
		if errors.Is(err, db.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, errorResponse(fmt.Errorf("blog not found")))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, updatedComment)
}

func (server *Server) deleteCommentByBlogID(ctx *gin.Context) {
	ID := ctx.Param("id")
	blogID, err := uuid.Parse(ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(fmt.Errorf("invalid blog id format")))
		return
	}

	commentID := ctx.Param("comment_id")
	commentUUID, err := uuid.Parse(commentID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(fmt.Errorf("invalid comment id format")))
		return
	}

	arg := db.DeleteCommentByBlogIDParams{
		ID:     commentUUID,
		BlogID: blogID,
	}

	err = server.store.DeleteCommentByBlogID(ctx, arg)
	if err != nil {
		if errors.Is(err, db.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
