package api

import (
	"context"
	"fmt"
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
	fmt.Println(param)
	usr, err := server.query.CreateUser(context.Background(), param.Username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, usr)
}
