package main

import (
	"fmt"
	"github.com/lyeka/gwt/router"

	"log"
	"net/http"
)

func main() {
	fmt.Println("hello, gwt")

	ginHandler := router.InitRouter()

	server := http.Server{
		Addr:    ":8080",
		Handler: ginHandler,
	}

	log.Fatal(server.ListenAndServe())

}
