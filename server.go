package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/configmap", ConfigMap)
	http.HandleFunc("/", Hello)
	http.ListenAndServe(":80", nil)
}

func ConfigMap(writer http.ResponseWriter, request *http.Request) {
	data, err := ioutil.ReadFile("myfamily/family.txt")
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	fmt.Fprintf(writer, "My family: %s", string(data))
}

func Hello(writer http.ResponseWriter, request *http.Request) {
	name := os.Getenv("NAME")
	age := os.Getenv("AGE")
	fmt.Fprintf(writer, "Hello %s, you are %s years old", name, age)
}
