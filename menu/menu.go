package menu

import (
	"authFLow/auth"
	"authFLow/data"
	"fmt"
)

func Menu(choce int) {
	var registered data.DataRegister
	switch choce {
	case 1:
		auth.AuthRegister()
		fmt.Println("Register success, press enter to back..")
		fmt.Scanln()
		
	case 2:
		auth.Login(registered)
	case 3:
		fmt.Println("Forget Password")
	case 0:
		fmt.Println("Exit Program")
	default:
		panic("Wrong input")
	}
}
