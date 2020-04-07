package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

var (
	host string = "0.0.0.0:8088"
)

func main() {
	http.HandleFunc("/", response)
	http.ListenAndServe(host, nil);
}

func response(writter http.ResponseWriter, request *http.Request) {
	targetURI, err :=ioutil.ReadAll(request.Body)
	if err != nil {
		panic(err)
	}
	body := httpPost(string(targetURI))
	writter.Write(body)
}

func httpPost(targetURI string) []byte {
	var (
		contentType string = "application/x-www-form-urlencoded"
	)
	resp, err := http.Post(targetURI, contentType, nil)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		//TODO handle error
		fmt.Println(err)
	}
	return body
}