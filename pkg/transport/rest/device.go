package rest

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) getDevices(c *gin.Context) {
	user_id, _ := c.Get(userCtx)

	devices, err := h.services.UserDevice.GetUserDevice(user_id.(int))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, devices)
}

type AddDeviceRequest struct {
	Uid string `json:"uid" binding:"required"`
}

func (h *Handler) addDevice(c *gin.Context) {
	user_id, _ := c.Get(userCtx)

	var deviceRequest AddDeviceRequest

	if err := c.BindJSON(&deviceRequest); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	device, err := h.services.UserDevice.AddUserDevice(user_id.(int), deviceRequest.Uid)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, device)
}

func (h *Handler) deleteDevice(c *gin.Context) {
	user_id, _ := c.Get(userCtx)
	device_id, _ := strconv.Atoi(c.Param("id"))

	err := h.services.UserDevice.DeleteUserDevice(user_id.(int), device_id)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]bool{
		"success": true,
	})
}
