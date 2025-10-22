package menu

import (
	"authFLow/auth"
	"authFLow/auth/home"
	"authFLow/data"
	"fmt"
)

var registered data.DataRegister

func Menu(choice int) bool {
	switch choice {
	case 1:
		registered = auth.AuthRegister()
		fmt.Println("Register success, press enter to back..")
		fmt.Scanln()
		return false

	case 2:
		if auth.Login(registered) {
			home.HomePage()
			return true 
		}
		return false 

	case 3:
		fmt.Println("Forget Password feature coming soon...")
		fmt.Println("Press enter to continue...")
		fmt.Scanln()
		return false

	default:
		panic("Wrong input")
	}
}