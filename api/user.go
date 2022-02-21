package api

import (
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
	Id int64 `uri:"id", binding:"required"`
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
