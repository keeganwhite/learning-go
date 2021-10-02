package main

import (
	"fmt"
	"os"
)

func main()  {
	if len(os.Args) < 2 {
		fmt.Println("No command provided")
		os.Exit(2)
	}
	fmt.Println(os.Args)
}
