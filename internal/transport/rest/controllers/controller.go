package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type response struct {
	Success bool
	Data    any
	Error   string
}

func out(w http.ResponseWriter, success bool, data any, err error) {

	var resp response

	if err == nil {
		resp.Success = true
		resp.Data = data
	} else {
		resp.Success = false
		resp.Error = err.Error()
	}

	json_bytes, _ := json.Marshal(resp)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(json_bytes))
}
