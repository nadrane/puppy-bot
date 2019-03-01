package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func puppyHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("url %+v", r.URL)
	queue := make(chan []byte)
	fmt.Println("calling GetPuppy")
	go GetPuppy(queue)
	w.Header().Set("Content-Type", "image/jpeg")
	puppyBuffer := <-queue
	w.Write(puppyBuffer)
}

func faviconHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}

func main() {
	http.HandleFunc("/favicon.icocalling", faviconHandler)
	http.HandleFunc("/puppy", puppyHandler)
	fmt.Printf("listening on http://localhost:%v \n", os.Getenv("PORT"))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", os.Getenv("PORT")), nil))
}
