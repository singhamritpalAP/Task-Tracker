package relationaldatabase

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"taskTracker/internal/models/tasktracker/dao"
)

// DbWrapper struct to hold the database connection
type DbWrapper struct {
	RelationalDatabase
	DB *gorm.DB
}

func NewDbWrapper() (*DbWrapper, error) {
	dbWrapper := &DbWrapper{}
	err := dbWrapper.init()
	if err != nil {
		return nil, err
	}
	return dbWrapper, nil
}

func (dbWrapper *DbWrapper) init() error {
	dbInstance, err := connectDB()
	if err != nil {
		log.Println("Error connecting to DB")
		return err
	}

	err = dbInstance.AutoMigrate(&dao.TaskTracker{}, &dao.User{})
	if err != nil {
		log.Println("Error auto migrating models")
		return err
	}
	dbWrapper.DB = dbInstance
	return nil
}

func connectDB() (*gorm.DB, error) {
	// Define connection string
	connString := "postgres://ap:31031999@localhost:5432/task_tracker_db" // fetch from config map

	// Connect to the database
	dbInstance, err := gorm.Open(postgres.Open(connString), &gorm.Config{})
	if err != nil {
		log.Println("Error connecting to DB")
		return nil, err
	}

	sqlDb, err := dbInstance.DB()
	if err != nil {
		log.Println("Error connecting to DB")
		return nil, err
	}

	err = sqlDb.Ping()
	if err != nil {
		log.Println("Error pinging DB")
		return nil, err
	}

	fmt.Println("Connected to PostgreSQL database!")
	return dbInstance, nil
}

// Create function to insert a new record
func (dbWrapper *DbWrapper) Create(model interface{}, values interface{}) error {
	log.Println("Create model", values)
	result := dbWrapper.DB.Model(model).Create(values)
	return result.Error
}

// Read function to find a record by TaskId
func (dbWrapper *DbWrapper) Read(model interface{}, id string) error {
	result := dbWrapper.DB.First(&model, id)
	return result.Error
}

// Update function to update an existing record
func (dbWrapper *DbWrapper) Update(model interface{}, value map[string]interface{}, condition string, args ...interface{}) error {
	result := dbWrapper.DB.Model(model).Where(condition, args...).Updates(value)
	return result.Error
}

// Delete function to remove a record by TaskId
func (dbWrapper *DbWrapper) Delete(model interface{}, id string, condition string, args ...interface{}) error {
	result := dbWrapper.DB.Where(condition, args...).Delete(model, id)
	return result.Error
}

// ReadAll function to retrieve all records of a specific type
func (dbWrapper *DbWrapper) ReadAll(models interface{}) error {
	result := dbWrapper.DB.Find(models)
	return result.Error
}
