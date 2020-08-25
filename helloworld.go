package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	log.Print("Hello world received a request.")
	target := "Welcome to CloudBase"
	fmt.Fprintf(w, "Hello, %s!\n", target)

	path := "/tmp/tcb/aaaa.log"
 	fp, fpErr := os.OpenFile(path, os.O_APPEND|os.O_CREATE, 0666)
	fp.WriteString("Welcome to CloudBase")

	if fpErr != nil {
		fp.WriteString("fp.WriteString failed.")
	}
	defer fp.Close()
}

func main() {
	log.Print("你好，恒娟.")

	http.HandleFunc("/", handler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
