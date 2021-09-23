package main

import "fmt"

func main() {
	doSomething()
	fmt.Println("8 + 2 =", addValues(8,2))
	length, total := addAllValues(8, 2, 3)
	fmt.Printf("8 + 2 + 3 = %v and the number of digits is %v\n", total, length)
	simba := Dog{"Bullmastiff", 60, "Woof"}
	simba.Speak()
}

// Dog is a struct
type Dog struct {
	Breed string
	Weight int
	Sound string
}

func (d Dog) Speak()  {
	fmt.Println(d.Sound)
}

func doSomething() {
	fmt.Println("This is a function")
}

func addValues(value1, value2 int)  int {
	return value1 + value2
}

func addAllValues(values ...int)  (int,int) {
	total := 0
	length := len(values)
	for _, v := range values {
		total += v
	}
	return length, total
}