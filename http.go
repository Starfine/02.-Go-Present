// +build OMIT
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", fHelloWorld)
	fmt.Println("Please open http://IP:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func fHelloWorld(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello worldddd")
	_ = r
	return
}
