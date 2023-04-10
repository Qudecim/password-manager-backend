package rest

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	service "github.com/qudecim/password-manager-backend/internal/services"
)

func DataGet(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	user, err := service.GetUser(r)
	if err != nil {
		out(w, false, nil, err)
		return
	}

	service.DataAdd(user, "")

	out(w, true, nil, err)
}
