package api

import (
	"fmt"

	"github.com/Thanhbinh1905/go-training-bank/internal/config"
	"github.com/Thanhbinh1905/go-training-bank/internal/service"
	"github.com/Thanhbinh1905/go-training-bank/internal/token"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type Server struct {
	cfg        config.Config
	service    service.Service
	tokenMaker token.Maker
	router     *gin.Engine
}

func NewServer(cfg config.Config, service service.Service) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(cfg.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	server := &Server{
		cfg:        cfg,
		service:    service,
		tokenMaker: tokenMaker,
	}

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}
	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	router.POST("/users", server.CreateUser)
	router.POST("/users/login", server.Login)

	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))

	authRoutes.POST("/accounts", server.CreateAccount)
	authRoutes.GET("/accounts/:id", server.GetAccount)
	authRoutes.GET("/accounts", server.ListAccounts)

	authRoutes.POST("/transfers", server.CreateTransfer)

	authRoutes.GET("/users/:username", server.GetUser)

	server.router = router
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
