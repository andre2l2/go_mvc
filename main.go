package main

import (
	"net/http"

	"github.com/andre2l2/go_mvc/routes"
)

func main() {	
	routes.RunRoutes()
	http.ListenAndServe(":8000", nil)
}
