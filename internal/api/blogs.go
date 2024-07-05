package api

import (
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
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

	ctx.JSON(http.StatusCreated, blog)

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

	// check if the UUID is nil (empty)
	if id == uuid.Nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(fmt.Errorf("invalid UUID: cannot be empty")))
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

type listBlogsRequest struct {
	PageID   int32 `form:"page_id" binding:"omitempty,min=1"`
	PageSize int32 `form:"page_size" binding:"omitempty,min=5,max=10"`
}

func (server *Server) listBlogs(ctx *gin.Context) {
	var req listBlogsRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// setting defaults incase they are empty in query
	if req.PageID == 0 {
		req.PageID = 1
	}
	if req.PageSize == 0 {
		req.PageSize = 10
	}

	arg := db.ListBlogsParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	blogs, err := server.store.ListBlogs(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, blogs)
}

type updateBlogRequest struct {
	ID          string `uri:"id" binding:"required,uuid"`
	Title       string `json:"title" binding:"omitempty,min=1"`
	Slug        string `json:"slug" binding:"omitempty,min=1"`
	Description string `json:"description" binding:"omitempty,min=1"`
	Body        string `json:"body" binding:"omitempty,min=1"`
	BannerImage string `json:"banner_image" binding:"omitempty,url"`
}

func (server *Server) updateBlog(ctx *gin.Context) {
	var req updateBlogRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	id, err := uuid.Parse(req.ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(fmt.Errorf("invalid UUID format")))
		return
	}

	arg := db.UpdateBlogParams{
		ID: id,
		Title: pgtype.Text{
			String: req.Title,
			Valid:  req.Title != "",
		},
		Slug: pgtype.Text{
			String: req.Slug,
			Valid:  req.Slug != "",
		},
		Description: pgtype.Text{
			String: req.Description,
			Valid:  req.Description != "",
		},
		Body: pgtype.Text{
			String: req.Body,
			Valid:  req.Body != "",
		},
		BannerImage: pgtype.Text{
			String: req.BannerImage,
			Valid:  req.BannerImage != "",
		},
	}

	blog, err := server.store.UpdateBlog(ctx, arg)
	if err != nil {
		if errors.Is(err, db.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, errorResponse(fmt.Errorf("blog not found")))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, blog)
}

type deleteBlogRequest struct {
	ID string `uri:"id" binding:"required"`
}

func (server *Server) deleteBlog(ctx *gin.Context) {
	var req deleteBlogRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	id, err := uuid.Parse(req.ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(fmt.Errorf("invalid UUID format")))
		return
	}

	err = server.store.DeleteBlog(ctx, id)
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
