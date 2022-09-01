package database

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name"     gorm:"not null"`
	Email    string `json:"email"    gorm:"not null, unique"`
	Password string `json:"password" gorm:"not null"`
}

type Call struct {
	gorm.Model
	Title    string
	Content  string
	Status   string
	UserID   int
	User     User
	Comments []Comment
}

type Comment struct {
	gorm.Model
	Content string
	CallID  int
	Call    Call
	UserID  int
	User    User
}

func Migrate(db *gorm.DB) {
	db.Debug().AutoMigrate(&User{}, &Call{}, &Comment{})
	// db.Debug().AutoMigrate(&User{})
}
