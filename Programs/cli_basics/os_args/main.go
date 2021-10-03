package main

import (
	"fmt"
	"os"
	"strings"
)
// Very messy way of doing this. Rather use flag sets.
func main()  {
	if len(os.Args) < 2 {
		fmt.Println("No command provided")
		os.Exit(2)
	}
	cmd := os.Args[1]
	switch cmd {
	case "greet":
		msg := "REMINDERS CLI - CLI BASICS"
		if len(os.Args) > 2 {
			f := strings.Split(os.Args[2], "=")
			if len(f) == 2 && f[0] == "--msg" {
				msg = f[1]
			}
		}
		fmt.Printf("Hello and welcome %s\n", msg)
	case "help":
		fmt.Println("some help message")
	default:
		fmt.Printf("unknown command %s\n", cmd)
	}
}
