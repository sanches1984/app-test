package main

import (
	"fmt"
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	log.Println("call hello")
	fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {
	log.Println("run test app")
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Println("error serving", err)
	}
}
