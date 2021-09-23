package main

import (
	"fmt"
)

func main()  {
	anInt := 42
	var p = &anInt
	fmt.Println("Value of p is: ", *p)
	fmt.Println("The memory address of p is: ", p)
}