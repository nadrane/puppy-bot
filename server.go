package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	puppyResp := GetPuppy()
	puppyBuffer, err := ioutil.ReadAll(puppyResp)
	if err != nil {
		fmt.Println("error with puppy buffer")
	}
	w.Header().Set("Content-Type", "image/jpeg")
	w.Write(puppyBuffer)
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("listening on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
