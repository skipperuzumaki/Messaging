package api

import (
	db "Messaging/db/sqlc"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type createMessageGroupRequest struct {
	FromUser int64 `json:"fromUser" binding:"required"`
	ToUser   int64 `json:"toUser" binding:"required"`
}

func (server *Server) createMessageGroup(ctx *gin.Context) {
	var param createMessageGroupRequest
	err := ctx.ShouldBindJSON(&param)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	var identifier string
	if param.FromUser > param.ToUser {
		identifier = strconv.Itoa(int(param.ToUser)) + "-" + strconv.Itoa(int(param.FromUser))
	} else {
		identifier = strconv.Itoa(int(param.FromUser)) + "-" + strconv.Itoa(int(param.ToUser))

	}
	par := db.CreateMessageGroupParams{
		Identifier: identifier,
		FromUser:   param.FromUser,
		ToUser:     param.ToUser,
	}
	usr, err := server.query.CreateMessageGroup(ctx, par)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, usr)
}

type getMessageGroupParam struct {
	Id int64 `uri:"id" binding:"required"`
}

func (server *Server) getMessageGroup(ctx *gin.Context) {
	var param getMessageGroupParam
	err := ctx.ShouldBindUri(&param)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	usr, err := server.query.GetMessageGroup(ctx, param.Id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, usr)
}

type GroupFromIdentifierRequest struct {
	ID string `json:"id" biding:"required"`
}

func (server *Server) getMessageGroupFromIdentifier(ctx *gin.Context) {
	var param GroupFromIdentifierRequest
	err := ctx.ShouldBindJSON(&param)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	usr, err := server.query.GetMessageGroupFromIdentifier(ctx, param.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, usr)
}

type deleteMessageGroupParam struct {
	Id int64 `uri:"id" binding:"required"`
}

// TODO: Look into Foriegn key Constraints
func (server *Server) deleteMessageGroup(ctx *gin.Context) {
	var param deleteMessageGroupParam
	err := ctx.ShouldBindUri(&param)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	err = server.query.DeleteMessageGroup(ctx, param.Id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, "{}")
}

type listMessageGroupRequest struct {
	Usr int64 `uri:"usr" biding:"required"`
}

func (server *Server) listMessageGroup(ctx *gin.Context) {
	var param listMessageGroupRequest
	err := ctx.ShouldBindUri(&param)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	usr, err := server.query.ListGroups(ctx, param.Usr)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, usr)
}
