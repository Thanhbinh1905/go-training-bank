package api

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/Thanhbinh1905/go-training-bank/internal/dto"
	"github.com/Thanhbinh1905/go-training-bank/internal/token"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgconn"
)

func (server *Server) CreateAccount(ctx *gin.Context) {
	var req *dto.CreateAccountRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.PasetoPayload)
	arg := dto.CreateAccountRequest{
		Currency: req.Currency,
	}

	account, err := server.service.CreateAccount(ctx, authPayload.Username, &arg)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			// log.Printf(
			// 	"\npg error: \ncode=%s \nmessage=%s \ndetail=%s \nhint=%s \ntable=%s \ncolumn=%s \nconstraint=%s\n",
			// 	pgErr.Code,
			// 	pgErr.Message,
			// 	pgErr.Detail,
			// 	pgErr.Hint,
			// 	pgErr.TableName,
			// 	pgErr.ColumnName,
			// 	pgErr.ConstraintName,
			// )
			if pgErr.Code == "23505" { // unique_violation
				ctx.JSON(http.StatusForbidden, errorResponse(errors.New("duplicate account")))
				return
			}
			if pgErr.Code == "23503" { // foreign_key_violation
				ctx.JSON(http.StatusForbidden, errorResponse(errors.New("invalid foreign key")))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, account)
}

func (server *Server) GetAccount(ctx *gin.Context) {
	var req *dto.GetAccountRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	account, err := server.service.GetAccount(ctx, req)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.PasetoPayload)
	if account.Owner != authPayload.Username {
		err := errors.New("account doesn't belong to the authenticated user")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, account)
}

func (server *Server) ListAccounts(ctx *gin.Context) {
	var req *dto.ListAccountsRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.PasetoPayload)
	accounts, err := server.service.ListAccounts(ctx, authPayload.Username, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, accounts)
}

func (server *Server) DeleteAccount(ctx *gin.Context) {
	var req *dto.DeleteAccountRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err := server.service.DeleteAccount(ctx, req)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, successResponse("Account deleted successfully"))
}
