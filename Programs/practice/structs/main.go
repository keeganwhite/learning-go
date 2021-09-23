package main

import (
	"fmt"
)

func main() {
	simba := Dog{"Bullmastiff", 60}
	fmt.Println(simba)
	fmt.Printf("%+v\n", simba)
	fmt.Printf("Breed: %v\nWeight: %v\n", simba.Breed, simba.Weight)
}

// Dog is a struct
type Dog struct {
	Breed string
	Weight int
}
