package controllers

import (
	"../../../services"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Message struct {
	Name string `json:"name"`
}

func GetAllDocuments(writer http.ResponseWriter, request *http.Request) {
	bolB, _ := json.Marshal(services.GetAllDocuments())
	fmt.Fprintf(writer, string(bolB))
}

func CreateDocument(writer http.ResponseWriter, request *http.Request) {

	b, err := ioutil.ReadAll(request.Body)
	if err != nil {
		fmt.Printf("Bad req")

		http.Error(writer, err.Error(), 500)
		return
	}

	var msg Message

	err = json.Unmarshal(b, &msg)
	if err != nil {

		http.Error(writer, err.Error(), 500)
		return
	}

	_, err = json.Marshal(msg)
	if err != nil {
		http.Error(writer, err.Error(), 500)
		return
	}

	fmt.Printf("%s", b)
	fmt.Fprintf(writer, services.CreateDocument(msg.Name))

}
