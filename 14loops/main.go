package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops in golang")

	days := []string{"Sunday", "Tuesday", "Wednesday", "Friday", "Saturday"}

	fmt.Println(days)

	//for d := 0; d < len(days); d++ {
	//	fmt.Println(days[d])
	//}

	//for i := range days {
	//	fmt.Println(days[i])
	//}

	//for _, day := range days {
	//	fmt.Println(day)
	//}

	//for index, day := range days {
	//	fmt.Printf("index is %v and value is %v\n", index, day)
	//}

	// Like while loop
	rougueValue := 1
	for rougueValue < 10 {
		if rougueValue == 2 {
			goto lco
		}
		if rougueValue == 5 {
			rougueValue++
			continue
		}
		fmt.Println("Value is:", rougueValue)
		rougueValue++
	}

lco:
	fmt.Println("Jumping at Learncodeonline.in")
}
