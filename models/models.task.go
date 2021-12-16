package models

import (
	"time"
)

type Task struct {
	TaskID uint   `gorm:"primary_key;auto_increment;not_null" json:"id"`
	UserID uint 	`gorm:"not_null" json:"user_id"`
	Name string 	`json:"name"`
	Description string	`json:"description"`
	Deadline time.Time  `json:"deadline"`
	Status bool  	`json:"status"`
	Color string 	`json:"color"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Tag struct {
	TagID uint    `gorm:"primary_key;auto_increment;not_null" json:"id"`
	UserID uint		`gorm:"not_null" json:"user_id"`
	Name string 	`json:"name"`
	Description string 	`json:"description"`
	Color string 	`json:"color"`
}

type TaskTag struct {
	TaskID uint		`gorm:"not_null" json:"task_id"`
	TagID uint 		`gorm:"not_null" json:"tag_id"`
}

type List struct {
	ListID uint   `gorm:"primary_key;auto_increment;not_null" json:"id"`
	Name string 	`json:"name"`
	Description string 	`json:"description"`
}

type TaskList struct {
	TaskID uint		`gorm:"not_null" json:"task_id"`
	ListID uint 	`gorm:"not_null" json:"list_id"`
}