package main

import (
	"io"
	"log"
	"net/http"
)

func someHandler(w http.ResponseWriter, r *http.Request) {
	client := &http.Client{
		Transport: &http.Transport{},
	}
	resp, err := client.Get("https://example.com")
	if err != nil {
		log.Fatal(err)
	}
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	_ = b
}
