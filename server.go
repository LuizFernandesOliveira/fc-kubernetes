package main

import "net/http"

func main() {
	http.HandleFunc("/", Hello)
	http.ListenAndServe(":80", nil)
}

func Hello(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("Hello World!!!"))
}
