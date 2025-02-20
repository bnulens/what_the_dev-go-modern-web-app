package main

import "fmt"

const CURRENT_YEAR = 2024

func main() {
	var age int = 20
	var agePointer *int = &age

	fmt.Printf("pre: %d \n", age)
	fmt.Printf("age: %d \n", agePointer)
	getYearOfBirth(agePointer)
	fmt.Printf("post: %d \n", age)
}

// getYearOfBirth changes the `age` variable into a year of birth
func getYearOfBirth(ageP *int) {
	*ageP = CURRENT_YEAR - *ageP
}
