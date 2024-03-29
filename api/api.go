package api

import (
	"net/http"
)

func Setup() {

	http.HandleFunc("GET /{videoName}/", playlist)
	http.HandleFunc("GET /{videoName}/{segment}", segment)
	http.HandleFunc("GET /items", fileList)
}