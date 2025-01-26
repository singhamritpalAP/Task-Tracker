package dao

import "time"

// TaskTracker represents task model for DB table todo set column limits(use varchar with size)
type TaskTracker struct {
	TaskId      string `gorm:"primaryKey"`
	Title       string `gorm:"not null"`
	Description string
	Status      string    `gorm:"default:'PENDING'"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
	// UserId      string todo future scope to allow RUD of task to user who created it
}
