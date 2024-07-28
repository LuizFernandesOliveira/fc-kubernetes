package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

var startedAt = time.Now()

func main() {
	http.HandleFunc("/health", Healthz)
	http.HandleFunc("/secret", Secret)
	http.HandleFunc("/configmap", ConfigMap)
	http.HandleFunc("/", Hello)
	http.ListenAndServe(":80", nil)
}

func Healthz(writer http.ResponseWriter, request *http.Request) {
	duration := time.Since(startedAt)
	if duration.Seconds() < 10 {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte(fmt.Sprintf("Duration: %v", duration.Seconds())))
	} else {
		writer.WriteHeader(http.StatusOK)
		writer.Write([]byte(fmt.Sprintf("Duration: %v", duration.Seconds())))
	}
}

func Secret(writer http.ResponseWriter, request *http.Request) {
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	fmt.Fprintf(writer, "User: %s, Password %s", user, password)
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
