package api

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/Thanhbinh1905/go-training-bank/internal/dto"
	"github.com/Thanhbinh1905/go-training-bank/pkg/util"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgconn"
)

func (server *Server) CreateUser(ctx *gin.Context) {
	var req *dto.CreateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	hasedPass, err := util.Hash(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	arg := &dto.CreateUserRequest{
		Username: req.Username,
		Password: hasedPass,
		FullName: req.FullName,
		Email:    req.Email,
	}

	user, err := server.service.CreateUser(ctx, arg)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" {
				ctx.JSON(http.StatusForbidden, errorResponse(errors.New("username/email is used")))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	rsp := dto.NewUserResponse(user)

	ctx.JSON(http.StatusOK, rsp)
}

func (server *Server) GetUser(ctx *gin.Context) {
	var req *dto.GetUserRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	user, err := server.service.GetUser(ctx, req)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rsp := dto.NewUserResponse(user)

	ctx.JSON(http.StatusOK, rsp)
}

func (server *Server) Login(ctx *gin.Context) {
	var req *dto.LoginUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	user, err := server.service.GetUser(ctx, &dto.GetUserRequest{Username: req.Username})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			ctx.JSON(http.StatusNotFound, errorResponse(errors.New("user not found")))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	err = util.CheckHash(req.Password, user.HashedPassword)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	accessToken, err := server.tokenMaker.CreateToken(req.Username, server.cfg.AccessTokenDuration)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rsp := &dto.LoginUserResponse{
		AccessToken: accessToken,
		User:        *dto.NewUserResponse(user),
	}
	ctx.JSON(http.StatusOK, rsp)
}
