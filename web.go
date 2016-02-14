package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

var _, _ = fmt.Print("ok")

func main() {
	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "ok")
	})
	log.Println("start server.")
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}
