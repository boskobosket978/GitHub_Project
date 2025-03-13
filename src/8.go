package main

import "fmt"

func main() {
	// Generate a random number between 1 and 100
	randNum := randInt(1, 100)

	// Check if the number is odd or even
	if randNum%2 == 0 {
		fmt.Println("The number is even.")
	} else {
		fmt.Println("The number is odd.")
	}
}

func randInt(min, max int) int {
	return min + (rand.Int()%(max-min))
}
