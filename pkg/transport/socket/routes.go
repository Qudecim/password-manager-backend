package socket

import (
	"encoding/json"

	"github.com/qudecim/password-manager-backend/pkg/models"
	"github.com/sirupsen/logrus"
)

type deviceMessage struct {
	Event  string        `json:"event"`
	Device models.Device `json:"device"`
}

func (c *Client) handleInitialization(message []byte) {
	var obj deviceMessage
	err := json.Unmarshal(message, &obj)
	if err != nil {
		logrus.Error(err)
	}

	device, err := c.services.Device.CreateDevice(obj.Device)
	if err != nil {
		logrus.Error(err)
	}

	c.responseInitialization(device)

}

func (c *Client) handleConnect(message []byte) {
	// Обработка сообщения типа "type1"
}

func (c *Client) handleConnectConfirm(message []byte) {
	// Обработка сообщения типа "type1"
}
