package main

import (
	"bufio"
	"fmt"
	"os"
)

const aConst int = 64

func main() {
	var aString string = "This is a weird way to assign strings, am I right?"
	fmt.Println("Hello from Go!\n")
	fmt.Println(aString)
	fmt.Printf("The variable's type is: %T\n", aString)

	var anInteger int = 337
	fmt.Println(anInteger)

	var defaultInteger int
	fmt.Println(defaultInteger)

	var implicitVariable = "\nImplicit string"
	fmt.Println(implicitVariable)
	fmt.Printf("The variable's type is: %T\n", implicitVariable)

	anotherString := "\nThis only works inside functions"
	fmt.Println(anotherString)
	fmt.Printf("The variable's type is: %T\n", anotherString)

	fmt.Println(aConst)
	fmt.Printf("The variable's type is: %T\n", aConst)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	input, _ := reader.ReadString('\n')
	fmt.Println("The input was:", input)
}
