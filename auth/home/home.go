package home

import (
	"authFLow/data"
	"fmt"
)

func HomePage(accountLogin data.DataRegister) {
	for {
		fmt.Println("\n--- Welcome to system bek bek bek bek ---")
		fmt.Println("Hello,", accountLogin.Email)
		fmt.Println(`
1. List All User
2. Logout
0. Exit
`)

		var choice string
		fmt.Print("Choose a menu: ")
		fmt.Scanln(&choice)

		switch choice {
		case "1":
			for {
				fmt.Println("\n--- List of All Users ---")
				for i, user := range data.Users {
					fmt.Printf("%d.\n", i+1)
					fmt.Printf("   Email   : %s\n", user.Email)
					fmt.Printf("   Password: %s\n\n", user.Password)
				}

				fmt.Println("0. Back")
				var back string
				fmt.Print("Choose: ")
				fmt.Scanln(&back)
				if back == "0" {
					break
				}
			}

		case "2":
			fmt.Println("Logging out...")

			for i, user := range data.Users {
				if user.Email == accountLogin.Email {
					data.Users = append(data.Users[:i], data.Users[i+1:]...)
					break
				}
			}

			fmt.Println("User logged out and removed from system.")
			fmt.Println("Returning to main menu...")
			return

		case "0":
			fmt.Println("Program exited.")
			panic("Exit program")

		default:
			fmt.Println("Invalid input.")
		}
	}
}
