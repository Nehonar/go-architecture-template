package main

import (
	"go-architecture-template/internal/util"
	"net/http"
)

func main() {
	logger := util.NewLogger()
	mux := http.NewServeMux()

	// TODO: register routes using handlers here

	logger.Println("Server reunning on :8080")
	http.ListenAndServe(":8080", mux)
}
