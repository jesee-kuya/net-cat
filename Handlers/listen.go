package handlers

import (
	"fmt"
	"net/http"
)

func ListenHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
}
