package main

import (
	"log"
	"net/http"

	"github.com/WtoJglg4/chat_app/internal/app"
)

func main() {
	//standart log
	log.SetFlags(log.Lshortfile)

	//ws server
	server := app.NewServer("/entry")
	go server.Listen()

	//http.Handle("/", http.FileServer(http.Dir("work/src/github.com/WtoJglg4/chat_app/webroot")))
	http.Handle("/", http.FileServer(http.Dir("webroot")))
	log.Fatal(http.ListenAndServe(":8080", nil))

}
