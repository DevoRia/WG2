package routes

import (
	controller "../controllers"
	"net/http"
)

func InitRoutes() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		switch request.Method {
		case http.MethodGet:
			controller.GetAllDocuments(writer, request)
		case http.MethodPost:
			controller.CreateDocument(writer, request)
		}
	})
}
