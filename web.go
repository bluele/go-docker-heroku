package main

import (
	"fmt"
	"log"
	"net/http"
)

var _, _ = fmt.Print("ok")

func main() {
	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "ok")
	})
	log.Println("start server.")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
