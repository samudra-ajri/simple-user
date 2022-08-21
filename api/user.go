package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/samudra-ajri/simple-user/db/sqlc"
)

type createUserRequest struct {
	Name string `json:"name" binding:"required"`
}

func (server *Server) createUser(ctx *gin.Context) {
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	name := req.Name

	user, err := server.store.CreateUser(ctx, name)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, user)
}

type displayUserRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) DisplayUser(ctx *gin.Context) {
	var req displayUserRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	user, err := server.store.DisplayUser(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, user)
}

type displayAllUsersRequest struct {
	PageID int32 `form:"page_id"`
	// PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size"`
	// PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) DisplayAllUsers(ctx *gin.Context) {
	var req displayAllUsersRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if req.PageID == 0 {
		req.PageID = 1
	}

	if req.PageSize == 0 {
		req.PageSize = 5
	}

	arg := db.DisplayAllUsersParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	users, err := server.store.DisplayAllUsers(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, users)
}
