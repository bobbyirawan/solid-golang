package singleresponsibilityprinciple

import (
	"fmt"
	"testing"
)

// UserManager handles user-related operations
type UserManager struct {
	Users []User
}

func (um *UserManager) AddUser(name, email string) {
	user := User{Name: name, Email: email}
	um.Users = append(um.Users, user)
	fmt.Println("User added:", user)
}

// DatabaseService handles persistence
type DatabaseService struct{}

func (ds *DatabaseService) SaveUsers(users []User) {
	fmt.Println("Saving users to the database...")
}

func TestGoodUserManagement(t *testing.T) {
	userManager := UserManager{}
	dbService := DatabaseService{}

	userManager.AddUser("John Doe", "john@example.com")
	dbService.SaveUsers(userManager.Users) // Separate responsibility for saving
}
