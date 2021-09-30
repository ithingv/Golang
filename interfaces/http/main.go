package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		log.Fatal("Error: ", err)
		os.Exit(1)
	}

	//fmt.Println(resp) // https://pkg.go.dev/net/http@go1.17.1#Response
	bs := make([]byte, 99999)
	resp.Body.Read(bs)
	fmt.Println(string(bs))

}
