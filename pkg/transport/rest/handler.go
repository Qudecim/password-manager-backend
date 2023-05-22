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
		auth.GET("/", h.getSecrets)
		auth.GET("/:id", h.getSecret)
		auth.POST("/", h.createSecret)
		auth.PUT("/:id", h.updateSecret)
		auth.DELETE("/:id", h.deleteSecret)
	}

	auth = router.Group("/device", h.userIdentity)
	{
		auth.GET("/", h.getDevices)
		auth.POST("/", h.addDevice)
		auth.DELETE("/:id", h.deleteDevice)
	}

	return router

}
