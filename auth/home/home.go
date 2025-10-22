package home

import (
	"authFLow/data"
	"fmt"
)

func HomePage() {
	var accoutLogin data.DataRegister

	fmt.Println("--- Welcome to sytem bek bek bek bek ---")
	fmt.Println("Hello,", accoutLogin.Email)
	fmt.Println(`
		1. List All User
		2. Logout

		0. Exit
		`)

	var choce string
	fmt.Print("Choose a menu:")
	fmt.Scanln(&choce)

	switch choce {
	case "1":
		fmt.Println("list user")
	case "2":
		fmt.Printf("log out")
	default:
		fmt.Print("tutup")
		break

	}

}
