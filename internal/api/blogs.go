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

type createAccountRequest struct {
	Title       string `json:"title" binding:"required"`
	Slug        string `json:"slug" binding:"required"`
	Description string `json:"description" binding:"required"`
	Body        string `json:"body" binding:"required"`
	BannerImage string `json:"banner_image" binding:"required"`
}

func (server *Server) createBlog(ctx *gin.Context) {
	var req createAccountRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		if err == io.EOF {
			ctx.JSON(http.StatusBadRequest, errorResponse(errors.New("request body is empty")))
		} else {
			ctx.JSON(http.StatusBadRequest, errorResponse(err))
		}
		return
	}

	arg := db.CreateBlogParams{
		Title:       req.Title,
		Slug:        req.Slug,
		Description: req.Description,
		Body:        req.Body,
		BannerImage: req.BannerImage,
	}

	blog, err := server.store.CreateBlog(ctx, arg)
	if err != nil {
		errCode := db.ErrorCode(err)
		if errCode == db.ForeignKeyViolation || errCode == db.UniqueViolation {
			ctx.JSON(http.StatusForbidden, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, blog)

}

type getBlogByIDRequest struct {
	ID string `uri:"id" binding:"required"`
}

func (server *Server) getBlogByID(ctx *gin.Context) {
	var req getBlogByIDRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	id, err := uuid.Parse(req.ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(fmt.Errorf("invalid UUID format")))
		return
	}

	blog, err := server.store.GetBlog(ctx, id)
	if err != nil {
		if errors.Is(err, db.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, blog)
}
