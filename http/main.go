package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	resp, err := http.DefaultClient.Get("https://google.com")
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	fmt.Println("Response status: ", resp.StatusCode)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(body))
}
