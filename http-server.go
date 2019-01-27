package main

import (
	"fmt"
	"log"
	"net/http"
)
func main() {
	http.HandleFunc("/", helloWorldHandler)
	err := http.ListenAndServe(":8080", nil)
        if err != nil {
           log.Fatal("error starting http server : ", err)
        }
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!\n")
}
