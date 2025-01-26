package userservice

import (
	"log"
	models "taskTracker/internal/models/tasktracker/tracker"
	"taskTracker/internal/models/utils"
	"taskTracker/internal/ports/tasktraceregress"
)

type UserService struct {
	db tasktraceregress.DbPort
}

func NewApplication(database tasktraceregress.DbPort) *UserService {
	return &UserService{
		db: database,
	}
}

// CreateUser for creating a new user in DB
func (userService *UserService) CreateUser(user models.User) error {
	log.Println("Creating task")
	userDetails := utils.ConvertUserToDao(user)
	log.Println("user: ", userDetails.Username)
	// todo store hash password, storing as plain text for now
	err := userService.db.GetUser().Create(userDetails)
	if err != nil {
		return err
	}
	log.Println("User created")
	return nil
}

// ValidateUser for validating user
// currently not implemented
func (userService *UserService) ValidateUser(user models.User) error {
	// todo validate user
	return nil
}
