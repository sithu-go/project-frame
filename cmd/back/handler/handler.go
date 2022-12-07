package handler

import (
	"h-pay/ds"
	"h-pay/middleware"
	"h-pay/repository"
	"h-pay/service"

	"github.com/gin-gonic/gin"
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

	// auth routes
	// authHandler := newAuthHandler(h)
	// authHandler.register()

	// // admin routes
	// adminHandler := newAdminHandler(h)
	// adminHandler.register()

	// // bank routes
	// bankHandler := newBankHandler(h)
	// bankHandler.register()
}
