package api

import (
	"github.com/gin-gonic/gin"
)

// Serves all HTTP requests for the banking service
type Server struct {
	store *db.store // allows to interact with the database when processing requests from the clients
	router *gin.Engine // helps us route each API request to the correct handler for processing
}

func NewServer(store *db.Store) *Server {
	server := &Server{store:store}
	router := gin.Default()

	router.POST("/accounts",  server.createAccount) // if you are passing multiple functions, the last one should be the handler, other others should be middlewares
	router.POST("/accounts/:id",  server.getAccount) // if you are passing multiple functions, the last one should be the handler, other others should be middlewares

	server.router = router
	return server
}


func (server *Server) Start(address string) error {
	return server.router.Run(address)
}