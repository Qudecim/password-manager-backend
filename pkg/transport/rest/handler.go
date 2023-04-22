package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/qudecim/password-manager-backend/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes(router *gin.Engine) *gin.Engine {

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	auth = router.Group("/secret", h.userIdentity)
	{
		auth.POST("/", h.createSecret)
	}

	return router

}
