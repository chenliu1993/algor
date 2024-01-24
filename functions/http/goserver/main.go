package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/hello", HelloGetHandler).Methods("GET")
	router.HandleFunc("/hello", HelloPostHandler).Methods("POST")

	log.Printf("begin...")
	http.ListenAndServe(":8090", router)
	log.Printf("end...")
}

func HelloGetHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
	w.WriteHeader(http.StatusOK)
}

func HelloPostHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("read body error: %v", err)
		w.Write([]byte("read body error"))
		w.WriteHeader(http.StatusBadRequest)
	}
	defer r.Body.Close()

	content, err := json.Marshal(body)
	if err != nil {
		log.Printf("unmarshal error: %v", err)
		w.Write([]byte("unmarshal error"))
		w.WriteHeader(http.StatusBadRequest)
	}

	file, err := os.OpenFile("hello.txt", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Printf("open file error: %v", err)
		w.Write([]byte("open file error"))
		w.WriteHeader(http.StatusBadRequest)
	}

	defer file.Close()

	_, err = file.Write(content)
	if err != nil {
		log.Printf("write file error: %v", err)
		w.Write([]byte("write file error"))
		w.WriteHeader(http.StatusBadRequest)
	}

	name := r.Header.Get("Name")

	responseBody := fmt.Sprintf("{'code': 0, 'message': 'Hi %s, success'}", name)
	w.Write([]byte(responseBody))

	w.WriteHeader(http.StatusOK)
}
