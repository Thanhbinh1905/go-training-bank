package api

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	db "github.com/Thanhbinh1905/go-training-bank/internal/db/sqlc"
	"github.com/Thanhbinh1905/go-training-bank/internal/dto"
	"github.com/Thanhbinh1905/go-training-bank/internal/token"
	"github.com/gin-gonic/gin"
)

func (server *Server) CreateTransfer(ctx *gin.Context) {
	var req *dto.CreateTransferRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	fromAccount, valid := server.validAccount(ctx, req.FromAccountID, req.Currency)
	if !valid {
		return
	}
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.PasetoPayload)
	if fromAccount.Owner != authPayload.Username {
		err := errors.New("the account doesn't belong to authenticated user")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	_, valid = server.validAccount(ctx, req.ToAccountID, req.Currency)
	if !valid {
		return
	}

	arg := dto.CreateTransferRequest{
		FromAccountID: req.FromAccountID,
		ToAccountID:   req.ToAccountID,
		Amount:        req.Amount,
		Currency:      req.Currency,
	}

	account, err := server.service.CreateTransfer(ctx, &arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, account)
}

func (server *Server) validAccount(ctx *gin.Context, acountID int64, currency string) (*db.Account, bool) {
	account, err := server.service.GetAccount(ctx, &dto.GetAccountRequest{ID: int64(acountID)})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return nil, false
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return nil, false
	}

	if account.Currency != currency {
		err := fmt.Errorf("account %d currency mismatch: %s (expected %s)", account.ID, account.Currency, currency)
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return nil, false
	}

	return account, true
}
