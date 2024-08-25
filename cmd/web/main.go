package main

import (
	"fmt"
	"net/http"

	"github.com/meetaayush/gotask/pkgs/config"
	"github.com/meetaayush/gotask/pkgs/handlers"
)

var app config.AppConfig

func main() {
	app.Environment = "Testing"

	repo := handlers.NewRepo(&app)
	handlers.NewHandler(repo)

	server := http.Server{
		Addr:    ":8080",
		Handler: routes(),
	}

	fmt.Println("Server is running on port 8080")
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
}
