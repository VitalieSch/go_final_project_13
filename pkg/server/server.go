package server

import (
	"fmt"
	"go1f/pkg/api"
	"net/http"
)

func Run() error {
	port := 7540
	http.Handle("/", http.FileServer(http.Dir("web")))

	api.Init()
	fmt.Println("Server is listening...")
	return http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
