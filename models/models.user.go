package models

type User struct {
	UserID   uint    	`gorm:"primary_key;auto_increment;not_null" json:"user_id"`
	Username string 	`gorm:"type:varchar(100);not_null" json:"username"`
	Password string 	`gorm:"type:varchar(100);not_null" json:"password"`
	Email		string 		`gorm:"type:varchar(100);not_null" json:"email"`
	FirstName string 	`gorm:"type:varchar(100);not_null" json:"first_name"`
	LastName string 	`gorm:"type:varchar(100);not_null" json:"last_name"`
	// RoleID int // check if the user is an admin
}