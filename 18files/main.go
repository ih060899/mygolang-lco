package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Welcome to files in golang.")
	content := "This needs to go in file - LeanCodeOnline.in"
	file, err := os.Create("./mylcogofile.txt")

	if err != nil {
		panic(err)
	}
	length, err := io.WriteString(file, content)

	if err != nil {
		panic(err)
	}
	fmt.Println("length is: ", length)
	defer file.Close()
}

func readFile(fileName string) {

}
