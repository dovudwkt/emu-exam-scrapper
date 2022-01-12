package main

import (
	"fmt"
	"log"
	"net/http"
)

var (
	// available exam periods in order as in the website
	periods = []string{"08:30", "12:30", "16:00"}

	// URL to be used for scrapping
	targetURL = "https://stdportal.emu.edu.tr/examlist.asp"
)

func server() {
	var port = 3001
	var addr = fmt.Sprintf(":%d", port)

	http.HandleFunc("/exams/import", importExamsHandler)
	http.HandleFunc("/exams/search", searchExamsHandler)

	log.Print("Starting server at port ", port)
	http.ListenAndServe(addr, nil)
}

func main() {
	server()
}
