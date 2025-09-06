package handler

import (
	"fmt"
	"net/http"
)

// ハンドラ、ルートが機能してるかのテスト用
func TestHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "success")
}
