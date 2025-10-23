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
		if user, ok := auth.Login(); ok {
			home.HomePage(user)
			return true
		}
		return false 

	case 3:
		auth.ForgetPassword()
		return false

	default:
		panic(fmt.Sprintf("Wrong input: %d", choice))
	}
}