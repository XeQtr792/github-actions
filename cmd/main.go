package main

import (
	api2 "github.com/github-actions/internal/api"
	"net/http"
)

func main()  {
	api := api2.Api{
		Client: http.DefaultClient,
	}

	api.GetStuff("https://jsonplaceholder.typicode.com/todos/1")
}