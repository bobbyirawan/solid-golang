package singleresponsibilityprinciple

import (
	"fmt"
	"testing"
)

// This struct is handling both user data and database operations
type UserService struct {
	Users []User
}

func (us *UserService) AddUser(name, email string) {
	user := User{Name: name, Email: email}
	us.Users = append(us.Users, user)
	fmt.Println("User added:", user)
}

func (us *UserService) SaveToDatabase() {
	fmt.Println("Saving users to the database...")
}

func TestUserManagement(t *testing.T) {
	userService := UserService{}

	userService.AddUser("John Doe", "john@example.com")
	userService.SaveToDatabase()
}
