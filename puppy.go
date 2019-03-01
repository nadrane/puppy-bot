package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type JSONResult struct {
	Status  string
	Message string
}

func GetPuppy(queue chan []byte) {
	resp, err := http.Get("https://dog.ceo/api/breeds/image/random")
	if err != nil {
		fmt.Printf("error getting puppy url: %v\n", err)
	}
	defer resp.Body.Close()

	dec := json.NewDecoder(resp.Body)
	var jsonBody JSONResult
	err = dec.Decode(&jsonBody)
	if err != nil {
		fmt.Println("error decoding json")
	}

	fmt.Println("Serving image", jsonBody.Message)

	resp, err = http.Get(jsonBody.Message)
	if err != nil {
		fmt.Printf("error puppy image: %v\n", err)
	}
	defer resp.Body.Close()

	puppyBuffer, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Printf("error reading puppy image stream: %v\n", err)
	}

	queue <- puppyBuffer
}
