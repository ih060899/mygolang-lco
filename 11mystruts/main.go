package main

import (
	"fmt"
)

func main() {
	fmt.Println("Structs in golang")
	// no inheritance in golang; No super or parent

	imran := User{"Imran", "imran@gmail.com", true, 24}
	fmt.Println(imran)
	fmt.Printf("Imran details are: %+v\n", imran)
	fmt.Printf("Name is %v and email is %v", imran.Name, imran.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
