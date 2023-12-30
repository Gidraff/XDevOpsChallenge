package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	rBody := make(map[string]interface{})
	rBody["message"] = "Welcome Page"
	rBody["status"] = http.StatusOK

	respBody, err := json.Marshal(rBody)
	if err != nil {
		log.Fatalf("Error marshaling json %s\n", err)
	}
	w.Write(respBody)
}

func itemsHandler(w http.ResponseWriter, r *http.Request) {
	rBody := make(map[string]interface{})
	rBody["message"] = "All Items Page"
	rBody["status"] = http.StatusOK

	respBody, err := json.Marshal(rBody)
	if err != nil {
		log.Fatalf("Error marshaling json %s\n", err)
	}
	w.Write(respBody)
}
func helloHandler(w http.ResponseWriter, r *http.Request) {
	rBody := make(map[string]interface{})
	rBody["message"] = "Hello Page!"
	rBody["status"] = http.StatusOK

	respBody, err := json.Marshal(rBody)
	if err != nil {
		log.Fatalf("Error marshaling json %s\n", err)
	}
	w.Write(respBody)
}

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/", indexHandler)
	router.HandleFunc("/hello", helloHandler)
	router.HandleFunc("/items", itemsHandler)

	log.Print("Server running on :3000 ...")
	if err := http.ListenAndServe(":3000", router); err != nil {
		log.Print("Server failed to start")
		os.Exit(1)
	}
}
