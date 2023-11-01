package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dovudwkt/emu-exam-scrapper/controllers"
)

func server() {
	var port = 3030
	var addr = fmt.Sprintf(":%d", port)

	http.HandleFunc("/exams/import", controllers.ImportExamsHandler)
	http.HandleFunc("/exams/search", controllers.SearchExamsHandler)

	log.Print("Starting server at port ", port)
	http.ListenAndServe(addr, nil)
}

func main() {
	server()
}
