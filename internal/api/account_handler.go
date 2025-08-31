package api

import (
	"net/http"

	"github.com/Thanhbinh1905/go-training-bank/internal/dto"
	"github.com/gin-gonic/gin"
)

func (server *Server) CreateAccount(ctx *gin.Context) {
	var req *dto.CreateAccountRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := dto.CreateAccountRequest{
		Owner:    req.Owner,
		Currency: req.Currency,
	}

	account, err := server.service.CreateAccount(ctx, &arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	ctx.JSON(http.StatusOK, account)
}
