package router

import (
	"net/http"

	"github.com/Fybrid/go-api-architecture/internal/http/app/handler"
)

func NewRouter() {
	// ルーティング
	// http.HandleFunc("/", handler.Handler)

	http.HandleFunc("/test", handler.TestHandler)
	http.HandleFunc("/api", handler.APIHandler)
}
