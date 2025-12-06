package response

import (
	"encoding/json"
	"net/http"
)

// APIResponse defines the standard JSON shape returned by the API.
type APIResponse struct {
	Status string      `json:"status"`
	Msg    string      `json:"msg"`
	Data   interface{} `json:"data,omitempty"`
}

// JSON sends a successfull JSON response.
func JSON(w http.ResponseWriter, status int, msg string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	json.NewEncoder(w).Encode(APIResponse{
		Status: "ok",
		Msg:    msg,
		Data:   data,
	})
}

// Error sends and error JSON response.
func Error(w http.ResponseWriter, status int, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	json.NewEncoder(w).Encode(APIResponse{
		Status: "error",
		Msg:    msg,
	})
}
