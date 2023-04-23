package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) getSecrets(c *gin.Context) {
	user_id, _ := c.Get(userCtx)

	secrets, err := h.services.Secret.GetAllSecrets(user_id.(int))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	c.JSON(http.StatusOK, secrets)

}

func (h *Handler) getOneSecret(c *gin.Context) {

}

func (h *Handler) createSecret(c *gin.Context) {
	id, _ := c.Get(userCtx)
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) updateSecret(c *gin.Context) {

}

func (h *Handler) deleteSecret(c *gin.Context) {

}
