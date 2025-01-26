package dao

import (
	"taskTracker/core/relationaldatabase"
	"taskTracker/internal/models/tasktracker/dao"
)

// User serves as an adapter for user-related database operations.
// It provides methods for managing user data through the relational database.
type User struct {
	wrapper relationaldatabase.RelationalDatabase // Database wrapper for executing queries
}

// NewUserAdapter initializes a new User adapter with the provided database wrapper.
// It returns a pointer to the User instance.
func NewUserAdapter(wrapper *relationaldatabase.DbWrapper) *User {
	return &User{wrapper: wrapper}
}

// Create inserts a new user into the database.
// It takes a User object containing user details and returns an error if the operation fails.
func (user *User) Create(userDetails dao.User) error {
	// Call the Create method on the wrapper to insert the userDetails into the database.
	return user.wrapper.Create(&dao.User{}, &userDetails)
}
