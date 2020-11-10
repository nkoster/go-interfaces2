package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	response, err := http.Get("https://google.com")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("%v\n", *response)
	// bodyBuf := make([]byte, 99999)
	// bytesRead, _ := response.Body.Read(bodyBuf)
	// fmt.Println(string(bodyBuf))
	// fmt.Println(bytesRead)
	io.Copy(os.Stdout, response.Body)
}
