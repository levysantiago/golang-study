package main

import (
	"fmt"
	"log"
	"net/http"
)

func main(){
	fileServer := http.FileServer(http.Dir("./static"))
	
	http.Handle("/", fileServer)

	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting server at port 3000\n")
	
	err := http.ListenAndServe(":3000", nil);

	if err != nil {
		log.Fatal(err)
	}
}

func formHandler(res http.ResponseWriter, req *http.Request ){
	err := req.ParseForm()
	if err != nil {
		fmt.Fprintf(res, "ParseForm() err: %v", err)
	}
	fmt.Fprintf(res, "POST request successful\n")
	name := req.FormValue("name")
	address := req.FormValue("address")

	fmt.Fprintf(res, "Name: %s\n", name)
	fmt.Fprintf(res, "Address: %s\n", address)
}

func helloHandler(res http.ResponseWriter, req *http.Request ){
	if req.URL.Path != "/hello" {
		http.Error(res, "404 not found", http.StatusNotFound)
		return
	}

	if req.Method != "GET" {
		http.Error(res, "method is not supported", http.StatusNotFound)
		return
	}

	fmt.Fprintf(res, "hello!")
}