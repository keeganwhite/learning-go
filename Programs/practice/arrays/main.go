package main

import "fmt"

func main() {
	var colours [3]string
	colours[0] = "Blue"
	colours[1] = "Green"
	colours[2] = "Red"
	fmt.Println(colours)
	fmt.Println(colours[0])

	var numbers = [5]int{1, 2, 3}
	fmt.Println(numbers)

	fmt.Println()
	fmt.Println("Number of colours: ", len(colours))
}
