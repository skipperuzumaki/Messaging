package api

import (
	db "Messaging/db/sqlc"

	"github.com/gin-gonic/gin"
)

type Server struct {
	query  *db.Queries
	router *gin.Engine
}

func NewServer(qry *db.Queries) *Server {
	server := &Server{query: qry}
	router := gin.Default()

	router.POST("/user", server.createUser)
	router.GET("/user/:id", server.getUser)
	router.PUT("/user", server.updateUser)
	router.DELETE("/user/:id", server.deleteUser)
	router.GET("user/all", server.listUser)

	router.POST("/messageGroup", server.createMessageGroup)
	router.GET("/messageGroup/:id", server.getMessageGroup)
	router.GET("/messageGroup/identifier", server.getMessageGroupFromIdentifier)
	router.DELETE("/messageGroup/:id", server.deleteMessageGroup)
	router.GET("messageGroup/all/:usr", server.listMessageGroup)

	router.POST("/message", server.createMessage)
	router.GET("/message/latest", server.getLatestUnreadMessage)
	router.PUT("/message/:grp", server.readMessageGroup)
	router.GET("message/all/", server.listMessage)

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
