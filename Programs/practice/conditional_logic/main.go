package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	theAnswer := 42
	var result string

	if theAnswer < 0 {
		result = "Less than 0"
	} else if theAnswer == 0 {
		result = "Equals zero"
	} else {
		result = "Greater than zero"
	}
	fmt.Println(result)

	if x := -42; x < 0 {
		result = "Less than 0"
	} else if x ==0 {
		result = "Equals zero"
	} else {
		result = "Greater than zero"
	}
	fmt.Println(result)
	fmt.Println()

	rand.Seed(time.Now().Unix())
	day := rand.Intn(7) + 1
	fmt.Println(day)

	switch day {
	case 1:
		result = "It's Sunday"
	case 2:
		result = "It's Monday"
	default:
		result = "It's not Sunday or Monday"
	}
	fmt.Println(result)
}
