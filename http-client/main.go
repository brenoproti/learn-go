package main

import (
	"io"
	"net/http"
	"time"
)

func main() {
	client := http.Client{Timeout: time.Millisecond}
	res, err := client.Get("http://google.com")
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	println(string(body))
}
