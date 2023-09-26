package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	req, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		panic(err)
	}

	res, err := io.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	req.Body.Close()

	fmt.Printf("%v\n", string(res))
}
