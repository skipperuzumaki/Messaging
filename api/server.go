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

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
