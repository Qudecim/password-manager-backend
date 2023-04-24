package rest

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/qudecim/password-manager-backend/pkg/models"
)

func (h *Handler) getSecrets(c *gin.Context) {
	user_id, _ := c.Get(userCtx)

	secrets, err := h.services.Secret.GetSecrets(user_id.(int))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, secrets)
}

func (h *Handler) getSecret(c *gin.Context) {
	user_id, _ := c.Get(userCtx)
	secret_id, _ := strconv.Atoi(c.Param("id"))

	secret, err := h.services.Secret.GetSecret(user_id.(int), secret_id)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, secret)
}

func (h *Handler) createSecret(c *gin.Context) {
	user_id, _ := c.Get(userCtx)

	var secretRequest models.Secret

	if err := c.BindJSON(&secretRequest); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	secret_id, err := h.services.Secret.CreateSecret(user_id.(int), secretRequest)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	secret, err := h.services.Secret.GetSecret(user_id.(int), secret_id)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, secret)
}

func (h *Handler) updateSecret(c *gin.Context) {
	user_id, _ := c.Get(userCtx)
	secret_id, _ := strconv.Atoi(c.Param("id"))

	var secretRequest models.Secret
	secretRequest.Id = secret_id
	if err := c.BindJSON(&secretRequest); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err := h.services.Secret.UpdateSecret(user_id.(int), secretRequest)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	secret, err := h.services.Secret.GetSecret(user_id.(int), secret_id)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, secret)
}

func (h *Handler) deleteSecret(c *gin.Context) {
	user_id, _ := c.Get(userCtx)
	secret_id, _ := strconv.Atoi(c.Param("id"))

	err := h.services.Secret.DeleteSecret(user_id.(int), secret_id)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]bool{
		"success": true,
	})
}
