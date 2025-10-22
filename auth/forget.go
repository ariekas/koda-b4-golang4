package auth

import (
	"authFLow/data"
	"fmt"
)

func ForgetPassword() {
	fmt.Println("--- Forget Password ---")

	var inputEmail string
	fmt.Print("Enter your registered email: ")
	fmt.Scanln(&inputEmail)

	found := false
	for i, user := range data.Users {
		if user.Email == inputEmail {
			found = true

			var newPassword, confirmPassword string
			for {
				fmt.Print("Enter your new password: ")
				fmt.Scanln(&newPassword)

				fmt.Print("Confirm your new password: ")
				fmt.Scanln(&confirmPassword)

				if newPassword == confirmPassword {
					data.Users[i].Password = newPassword
					data.Users[i].ConfirmPassword = confirmPassword
					fmt.Println("Password changed successfully!")
					break
				} else {
					fmt.Println("Passwords do not match. Try again.")
				}
			}
			break
		}
	}

	if !found {
		fmt.Println("Email not found. Please register first.")
	}

	fmt.Println("Press Enter to return to main menu...")
	fmt.Scanln()
}
