package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://lco.dev"

func main() {
	fmt.Println("LCO Web Request")

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Response is of type: %T\n", response)

	defer response.Body.Close()
	datatypes, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	content := string(datatypes)
	fmt.Println(content)
}
