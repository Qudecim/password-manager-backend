package rest

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/qudecim/password-manager-backend/internal/models"
	"github.com/qudecim/password-manager-backend/internal/models/mysql"
	service "github.com/qudecim/password-manager-backend/internal/services"
)

func DeviceList(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	user, err := service.GetUser(r)
	if err != nil {
		out(w, false, nil, err)
		return
	}

	devices, err := service.DeviceList(user)
	if err != nil {
		out(w, false, nil, err)
		return
	}

	out(w, true, devices, err)
}

func DeviceAdd(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	user, err := service.GetUser(r)
	if err != nil {
		out(w, false, nil, err)
		return
	}

	var deviceRequest models.Device

	err = json.NewDecoder(r.Body).Decode(&deviceRequest)
	if err != nil {
		out(w, false, nil, err)
		return
	}

	var device models.Device

	// Нужно проверить есть ли устройство с таким UUID на сокетах в текущий момент
	// Если есть то берем информацию из соектов и записываем девайс в БД
	// отвечаем что все ок

	device.UserId = user.ID

	err = mysql.DeviceAdd(&device)
	if err != nil {
		out(w, false, nil, err)
		return
	}

	out(w, true, nil, nil)
}

func DeviceDelete(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	user, err := service.GetUser(r)
	if err != nil {
		out(w, false, nil, err)
		return
	}

	var deviceRequest models.Device

	err = json.NewDecoder(r.Body).Decode(&deviceRequest)
	if err != nil {
		out(w, false, nil, err)
		return
	}

	device, err := mysql.DeviceGet(user, deviceRequest.Uuid)
	if err != nil {
		out(w, false, nil, err)
		return
	}

	err = service.DeviceDelete(device)
	if err != nil {
		out(w, false, nil, err)
		return
	}

	out(w, true, nil, nil)
}

func SendToDevice(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	out(w, true, nil, nil)
}
