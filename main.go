package main

import (
	documentRouter "./api/v1/routes"
	"log"
	"net/http"
)

func main() {
	documentRouter.InitRoutes()
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
