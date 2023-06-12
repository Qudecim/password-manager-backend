package socket

import (
	"encoding/json"

	"github.com/qudecim/password-manager-backend/pkg/models"
	"github.com/sirupsen/logrus"
)

type response struct {
	event string
	data  any
}

func (c *Client) responseInitialization(device models.Device) {

	sendData, err := json.Marshal(device)
	if err != nil {
		logrus.Error(err)
	}

	c.send <- sendData
}
