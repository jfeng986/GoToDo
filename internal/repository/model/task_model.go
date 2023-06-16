package model

import (
	"gorm.io/gorm"
)

type Status string

const (
	ToDo       Status = "Todo"
	InProgress Status = "In Progress"
	Done       Status = "Done"
)

type Task struct {
	gorm.Model
	User    User   `gorm:"foreignKey:UserID"`
	UserID  uint   `gorm:"not null"`
	Title   string `gorm:"type:varchar(50);not null;index"`
	Content string `gorm:"type:longtext;not null"`
	Status  Status `gorm:"not null"`
}
