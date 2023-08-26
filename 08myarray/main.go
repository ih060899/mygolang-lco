package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in golangs")

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Tomato"
	fruitList[3] = "Peach"

	fmt.Println("Fruit list is : ", fruitList)
	fmt.Println("Array length: ", len(fruitList))

	var vegList = [5]string{"Potato", "Beans", "Mushroom"}
	fmt.Println("Veg List is: ", vegList)
	fmt.Println("Veg List length: ", len(vegList))
}
