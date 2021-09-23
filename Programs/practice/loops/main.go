package main

import (
	"fmt"
)

func main() {
	colours := []string{"Red", "Green", "Blue"}
	fmt.Println(colours)

	for i := 0; i < len(colours); i++ {
		fmt.Println(colours[i])
	}

	for i := range colours {
		fmt.Println(colours[i])
	}

	for _, colour := range  colours {
		fmt.Println(colour)
	}

	fmt.Println()
	// Make a 'while loop'

	value := 1
	for value < 10 {
		fmt.Println("Value", value)
		value++
	}
}
