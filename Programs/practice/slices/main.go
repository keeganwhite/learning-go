package main

import (
	"fmt"
	"sort"
)

func main() {
	var colours = []string{"red", "green", "blue"}
	fmt.Println(colours)
	colours = append(colours, "purple")
	fmt.Println(colours)

	colours = append(colours[1:len(colours)])
	fmt.Println(colours)

	numbers := make([]int, 5)
	numbers[0] = 100
	numbers[1] = 101
	numbers[2] = 102
	numbers[3] = 103
	numbers[4] = 104
	fmt.Println(numbers)
	numbers = append(numbers, 15)
	fmt.Println(numbers)
	sort.Ints(numbers)
	fmt.Println(numbers)
}
