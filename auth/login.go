package auth

import (
	"authFLow/data"
	"fmt"
)

func Login(user data.DataRegister) {
	fmt.Println("--- Login ---")

	var inputEmail, inputPassword string

	fmt.Print("Enter your email :")
	fmt.Scan(&inputEmail)

	fmt.Print("Enter your password :")
	fmt.Scan(&inputPassword)

	if user.Email == inputEmail && user.Password == inputPassword {
		fmt.Println("Login success, press enter to back.. ")
	} else {
		fmt.Println("Login gagal.. ")
		return
	}
}
