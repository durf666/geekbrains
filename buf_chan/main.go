package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func getRequest(url string, ch chan<- string) {
	resp, err := http.Get("https://google.it")
	if err != nil {
		log.Fatal(err)
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	bodyString := string(bodyBytes)
	ch <- bodyString
}

func main() {
	fmt.Println(mirroredQuery())
}
func mirroredQuery() string {
	responses := make(chan string, 3)
	go getRequest("https://google.cn", responses)
	go getRequest("https://ya.ru", responses)
	go getRequest("https://google.at", responses)

	return <-responses

}
