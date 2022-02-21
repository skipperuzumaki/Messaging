package api

import (
	db "Messaging/db/sqlc"
	"net/http"

	"github.com/gin-gonic/gin"
)

type createUserRequest struct {
	Username string `json:"username" binding:"required"`
}

func (server *Server) createUser(ctx *gin.Context) {
	var param createUserRequest
	err := ctx.ShouldBindJSON(&param)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	usr, err := server.query.CreateUser(ctx, param.Username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, usr)
}

type getUserParam struct {
	Id int64 `uri:"id" binding:"required"`
}

func (server *Server) getUser(ctx *gin.Context) {
	var param getUserParam
	err := ctx.ShouldBindUri(&param)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	usr, err := server.query.GetUser(ctx, param.Id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, usr)
}

type updateUserRequest struct {
	ID       int64  `json:"id" biding:"required"`
	Username string `json:"username" binding:"required"`
}

func (server *Server) updateUser(ctx *gin.Context) {
	var param updateUserRequest
	err := ctx.ShouldBindJSON(&param)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	par := db.UpdateUserParams{
		ID:       param.ID,
		Username: param.Username,
	}
	usr, err := server.query.UpdateUser(ctx, par)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, usr)
}

type deleteUserParam struct {
	Id int64 `uri:"id" binding:"required"`
}

// TODO: Look into Foriegn key Constraints
func (server *Server) deleteUser(ctx *gin.Context) {
	var param deleteUserParam
	err := ctx.ShouldBindUri(&param)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	err = server.query.DeleteUser(ctx, param.Id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, "{}")
}

type listUserRequest struct {
	Limit  int32 `json:"limit" biding:"required"`
	Offset int32 `json:"offset" binding:"required"`
}

func (server *Server) listUser(ctx *gin.Context) {
	var param listUserRequest
	err := ctx.ShouldBindJSON(&param)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	par := db.ListUsersParams{
		Limit:  param.Limit,
		Offset: param.Offset,
	}
	usr, err := server.query.ListUsers(ctx, par)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, usr)
}
