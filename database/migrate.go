package database

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name"     gorm:"not null"`
	Email    string `json:"email"    gorm:"not null;unique"`
	Password string `json:"password" gorm:"not null"`
}

type Call struct {
	gorm.Model
	Title   string `json:"title"   gorm:"not null"`
	Content string `json:"content" gorm:"not null"`
	Status  string `json:"status"  gorm:"not null"`
	UserID  int    `json:"user_id" gorm:"not null"`
	//User    User   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:UserID;references:ID"`
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
}
