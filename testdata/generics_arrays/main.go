package main

import (
	"net/http"

	"github.com/arkannsk/swag/testdata/generics_basic/api"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @host localhost:4000
// @basePath /api
func main() {
	http.HandleFunc("/posts/", api.GetPost)
	http.ListenAndServe(":8080", nil)
}
