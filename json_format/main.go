package main

import (
	"encoding/json"
	"os"
)

func main() {
	file, err := os.Create("test.json")
	if err != nil {
		panic(err)
	}
	encorder := json.NewEncoder(file)
	encorder.SetIndent("", " ")
	encorder.Encode(map[string]string{
		"example": "ecording/json",
		"hello":   "world",
	})
}
