package auth

import (
	"authFLow/auth/utils"
	"authFLow/data"
	"fmt"
	"strings"
)

func AuthRegister() data.DataRegister {
	fmt.Println("--- Register ---")

	var inputFirstName, inputLastName, inputEmail, inputPassword, inputPasswordConfirm string
	var users data.DataRegister

	fmt.Print("What is your first name: ")
	fmt.Scanln(&inputFirstName)

	fmt.Print("What is your last name: ")
	fmt.Scanln(&inputLastName)

	fmt.Print("What is your email: ")
	fmt.Scanln(&inputEmail)

	for {
		fmt.Print("Enter a strong password: ")
		fmt.Scanln(&inputPassword)

		fmt.Print("Confirm your password: ")
		fmt.Scanln(&inputPasswordConfirm)

		if inputPassword == inputPasswordConfirm {
			fmt.Println("Register successful!")

			fmt.Println("\n \n Is it true?")
			fmt.Println("Firt Name", inputFirstName)
			fmt.Println("Last Name", inputLastName)
			fmt.Println("Email Name", inputEmail)
			fmt.Print("Continue (y/n)")

			var confirm string
			fmt.Scanln(&confirm)
			confirm = strings.ToLower(confirm)

			switch confirm {
			case "y":
				hashedPassword := utils.HashPassword(inputPassword)
				users = data.DataRegister{
					FirstName:       inputFirstName,
					LastName:        inputLastName,
					Email:           inputEmail,
					Password:        hashedPassword,
					ConfirmPassword: hashedPassword,
				}
				data.Users = append(data.Users, users)
				return users
			case "n":
				return AuthRegister()
			default:
				fmt.Print("gagal")
			}
			break
		} else {
			fmt.Println("Wrong confirm password, press Enter to try again...")
			fmt.Scanln()
		}
	}

	return users
}
