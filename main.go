package main

import (
	"fmt"
	"log"
	"net/http"
)

func server() {
	var port = 3030
	var addr = fmt.Sprintf(":%d", port)

	http.HandleFunc("/exams/import", importExamsHandler)
	http.HandleFunc("/exams/search", searchExamsHandler)

	log.Print("Starting server at port ", port)
	http.ListenAndServe(addr, nil)
}

func main() {
	server()
}
