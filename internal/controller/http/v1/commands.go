package v1

import (
	"PostgresProfessionalTestTask/internal/entity"
	"PostgresProfessionalTestTask/internal/service"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

type walletRoutes struct {
	w service.CommandsOperations
	l *slog.Logger
}

func newWalletRoutes(handler *gin.RouterGroup, w service.CommandsOperations, l *slog.Logger) {
	r := &walletRoutes{w, l}

	h := handler.Group("/wallet")
	{
		h.POST("/create", r.createNewWallet)
	}
}

func (r *walletRoutes) createNewWallet(c *gin.Context) {
	command := new(entity.Command)

	wallet, err := r.w.CreateCommand(c.Request.Context(), command)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, wallet)
}
