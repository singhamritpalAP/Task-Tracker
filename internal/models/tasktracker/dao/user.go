package dao

// User represents the user model for DB table todo set column limits(use varchar with size)
type User struct {
	UserId   uint   `gorm:"primaryKey" json:"user_id"`
	Username string `gorm:"unique;not null" json:"username"`
	Password string `json:"password"`
}
