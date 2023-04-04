package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func out(w http.ResponseWriter, success bool, data any, err error) {
	//http.Error(w, err.Error(), http.StatusBadRequest)
	json_bytes, _ := json.Marshal(data)
	fmt.Fprintf(w, string(json_bytes))
}
