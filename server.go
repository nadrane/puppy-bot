package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
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
	fmt.Printf("listening on http://localhost:%v \n", os.Getenv("PORT"))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", os.Getenv("PORT")), nil))
}
