package main

import (
	"fmt"
	"log"
	"net/http"
)

func handleRootGet(w http.ResponseWriter, r *http.Request) {

}

func main() {
	fmt.Println("Demo: Http Server")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("We have hit / path " + r.Method))
		switch r.Method {
		case "GET":
			handleRootGet(w, r)
		case "POST":
		case "PUT":
		case "DELETE":
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)

		}
		w.WriteHeader(http.StatusOK)
	})
	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("We have hit /test path " + r.Method))
	})
	http.HandleFunc("/test/case1", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("We have hit /test/case1"))
	})
	//go routine...
	go createConfigServer()
	err := http.ListenAndServe(":1234", nil)
	if err != nil {
		log.Fatalln("Failed to start server")
	}
}

func createConfigServer() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hitting root of config server"))
		w.WriteHeader(http.StatusOK)
	})

	err := http.ListenAndServe(":4321", mux)
	if err != nil {
		fmt.Println(err)
		// TRL, Alarm
	}

}
