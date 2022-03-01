package api

import (
	db "Messaging/db/sqlc"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type createMessageRequest struct {
	Message  string `json:"message" binding:"required"`
	FromUser int64  `json:"fromUser" binding:"required"`
	ToUser   int64  `json:"toUser" binding:"required"`
}

func (server *Server) createMessage(ctx *gin.Context) {
	var param createMessageRequest
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
	par := db.CreateMessageParams{
		Group:    identifier,
		Message:  param.Message,
		SentFrom: param.FromUser,
		SentTo:   param.ToUser,
	}
	usr, err := server.query.CreateMessage(ctx, par)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, usr)
}

type FromIdentifierRequest struct {
	ID int64 `json:"id" biding:"required"`
}

func (server *Server) getLatestUnreadMessage(ctx *gin.Context) {
	var param FromIdentifierRequest
	err := ctx.ShouldBindJSON(&param)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	usr, err := server.query.GetLatestUnreadMessage(ctx, param.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, usr)
}

type readMessageGroupParam struct {
	Grp string `uri:"grp" binding:"required"`
}

func (server *Server) readMessageGroup(ctx *gin.Context) {
	var param readMessageGroupParam
	err := ctx.ShouldBindUri(&param)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	err = server.query.ReadMessageGroup(ctx, param.Grp)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, "{}")
}

type listMessageRequest struct {
	Grp    string `json:"group" biding:"required"`
	Limit  int32  `json:"limit" biding:"required"`
	Offset int32  `json:"offset" biding:"required"`
}

func (server *Server) listMessage(ctx *gin.Context) {
	var param listMessageRequest
	err := ctx.ShouldBindJSON(&param)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	par := db.RetrieveAllParams{
		Group:  param.Grp,
		Limit:  param.Limit,
		Offset: param.Offset,
	}
	usr, err := server.query.RetrieveAll(ctx, par)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, usr)
}
