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

	fmt.Println("--- Welcome to system ---")
	fmt.Println(`
	1. Register
	2. Login
	3. Forget Password
	0. Exit
	`)
	
	var choce int
	fmt.Print("Choose a menu: ")
	fmt.Scan(&choce)
	menu.Menu(choce)

}


