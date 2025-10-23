package data

import "fmt"

type DataRegister struct {
	FirstName       string
	LastName        string
	Email           string
	Password        string
	ConfirmPassword string
}

func (d DataRegister) GetFullName() string {
	return fmt.Sprintf("Fullname: %s %s", d.FirstName, d.LastName)
}

type UserService interface {
	GetAllUsers() []DataRegister
}

type UserServiceImpl struct{}

func (s *UserServiceImpl) GetAllUsers() []DataRegister {
	return Users
}

var Users []DataRegister
