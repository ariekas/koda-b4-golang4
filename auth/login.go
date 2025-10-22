package auth

import (
	"authFLow/data"
	"fmt"
)

func Login() (data.DataRegister, bool) {
	for {
		fmt.Println("--- Login ---")

		var inputEmail, inputPassword string

		fmt.Print("Enter your email: ")
		fmt.Scanln(&inputEmail)

		fmt.Print("Enter your password: ")
		fmt.Scanln(&inputPassword)

		for _, user := range data.Users {
			if user.Email == inputEmail && user.Password == inputPassword {
				fmt.Println("Login success, press enter to back..")
				fmt.Scanln()
				return user, true
			}
		}

		fmt.Println("\nWrong email or password, press enter to restart.")
		fmt.Println("0. Home")

		var choice string
		fmt.Scanln(&choice)

		if choice == "0" {
			fmt.Println("Returning to main menu...")
			return data.DataRegister{}, false
		}
	}
}
