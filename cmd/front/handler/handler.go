package handler

import (
	"h-pay/ds"
	"h-pay/middleware"
	"h-pay/repository"
	"h-pay/service"
	"h-pay/utils"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type Handler struct {
	R    *gin.Engine
	repo *repository.Repository
}

type HConfig struct {
	R  *gin.Engine
	DS *ds.DataSource
}

func NewHandler(c *HConfig) *Handler {
	svc := service.NewService()
	repo := repository.NewRepository(c.DS, svc)
	return &Handler{
		R:    c.R,
		repo: repo,
	}
}

func (h *Handler) Register() {
	h.R.Use(middleware.Cors())

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("checkphrase", utils.Checkphrase)
	}

	// auth routes
	// 	authHandler := newAuthHandler(h)
	// 	authHandler.register()

	// // wallet routes
	// walletHandler := newWalletHandler(h)
	// walletHandler.register()
}
