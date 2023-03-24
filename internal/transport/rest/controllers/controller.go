package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func out(w http.ResponseWriter, data any) {
	json_bytes, _ := json.Marshal(data)
	fmt.Fprintf(w, string(json_bytes))
}
