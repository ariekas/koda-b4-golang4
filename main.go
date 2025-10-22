package main

import (
	"authFLow/menu"
	"fmt"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Error:", r)
		}
	}()

	for {
		fmt.Println("--- Welcome to system ---")
		fmt.Println(`
		1. Register
		2. Login
		3. Forget Password
		0. Exit
		`)

		var choice int
		fmt.Print("Choose a menu: ")
		fmt.Scanln(&choice)

		if choice == 0 {
			fmt.Println("Exit Program")
			break
		}

		if menu.Menu(choice) {
			break
		}
	}
}