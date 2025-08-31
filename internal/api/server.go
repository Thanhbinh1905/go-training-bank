package api

import (
	"github.com/Thanhbinh1905/go-training-bank/internal/service"
	"github.com/gin-gonic/gin"
)

type Server struct {
	service service.Service
	router  *gin.Engine
}

func NewServer(service service.Service) *Server {
	server := &Server{
		service: service,
	}
	router := gin.Default()

	router.POST("/accounts", server.CreateAccount)
	router.GET("/accounts/:id", server.GetAccount)
	router.GET("/accounts", server.ListAccounts)

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func successResponse(message string) gin.H {
	return gin.H{"message": message}
}
