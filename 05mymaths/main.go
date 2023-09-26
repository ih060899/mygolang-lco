package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
	fmt.Println("Welcome to maths in golang...")

	//var myNumberOne int = 2
	//var myNumberTwo float64 = 4

	//fmt.Println("The sum is: ", myNumberOne+int(myNumberTwo))

	// Random number
	//rand.Seed(time.Now().UnixNano())
	//fmt.Println(rand.Intn(5) + 1)

	// random from crypto

	myRandomNumber, _ := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println(myRandomNumber)

}
