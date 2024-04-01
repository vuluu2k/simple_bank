package api

import (
	"github.com/gin-gonic/gin"

	db "github.com/vuluu2k/simple_bank/db/sqlc"
)

// Server serves http requests for our banking services.
type Server struct {
	store  db.Store
	router *gin.Engine
}

// NewServer create a http server and setup routing.
func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccount)
	server.router = router

	return server
}

// Start runs the http server on a specific address.
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
