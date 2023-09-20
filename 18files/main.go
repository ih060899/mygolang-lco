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

	checkNilErr(err)
	length, err := io.WriteString(file, content)

	checkNilErr(err)
	fmt.Println("length is: ", length)
	defer file.Close()
	readFile("./mylcogofile.txt")
}

func readFile(fileName string) {
	//databyte, err := ioutil.ReadFile(fileName)
	databyte, err := os.ReadFile(fileName)
	checkNilErr(err)
	fmt.Println("Text data inside the file is \n", string(databyte))
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
