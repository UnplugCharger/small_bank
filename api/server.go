package api

import (
	db "github.com/UnplugCharger/small_bank/db/sqlc"
	"github.com/gin-gonic/gin"
)

// Server serves http requests for our banking service
type Server struct {
	store  *db.Store
	router *gin.Engine
}

//New server creates a new http server and set up routing

func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccount)

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{
		"error": err.Error(),
	}
}
