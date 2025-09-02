package api

import (
	"os"
	"testing"
	"time"

	"github.com/Thanhbinh1905/go-training-bank/internal/config"
	db "github.com/Thanhbinh1905/go-training-bank/internal/db/sqlc"
	"github.com/Thanhbinh1905/go-training-bank/internal/service"
	"github.com/Thanhbinh1905/go-training-bank/pkg/util"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func NewTestServer(t *testing.T, store db.Store) *Server {
	config := config.Config{
		TokenSymmetricKey:   util.RandomString(32),
		AccessTokenDuration: time.Minute,
	}
	svc := service.NewService(store)
	server, err := NewServer(config, svc)
	require.NoError(t, err)

	return server
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)

	os.Exit(m.Run())
}
