package menu

import "fmt"

func Menu(choce int) {
	switch choce {
	case 1:
		fmt.Println("Register")
	case 2:
		fmt.Println("Login")
	case 3:
		fmt.Println("Forget Password")
	case 0:
		fmt.Println("Exit Program")
	default:
		panic("Wrong input")
	}
}