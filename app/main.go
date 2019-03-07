package main

import (
	"fmt"

	"github.com/gravetii/gotest/http"
)

func main() {
	resp := http.Response{Message: "default message", Code: 200}
	fmt.Println(resp.Message)
	fmt.Println(resp.Code)

	req := http.Request{Header: "This is the header!", Body: "body"}
	fmt.Println(req.Header)
	fmt.Println(req.Body)
}
