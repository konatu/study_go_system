package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	type Response struct {
		Status string `json:"status"`
		Data   string `json:"data"`
	}

	response := Response{
		Status: "ok",
		Data:   "success",
	}

	json, _ := json.Marshal(response)

	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write(json)
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8088", nil))
}
