package main

import "fmt"

const LoginToken string = "dkajek" // Public

func main() {
	var username string = "Imran"
	fmt.Println(username)
	fmt.Printf("Vaiable is of type: %T \n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallval uint8 = 255
	fmt.Println(smallval)
	fmt.Printf("Variable is of type: %T \n", smallval)

	var smallFloat float64 = 255.345354080988
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of tyoe: %T \n", smallFloat)

	// default values and some alias
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T \n", anotherVariable)

	// implicit type
	var website = "LearnCodeOnline.in"
	fmt.Println(website)

	// no var style

	numberOfUsers := 3987493.0
	fmt.Println(numberOfUsers)
	fmt.Printf("Variable is of type: %T \n", numberOfUsers)

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type: %T \n", LoginToken)
}
