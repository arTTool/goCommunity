package model

type User struct {
	ID       string `gorm:"primaryKey;autoIncrement"`
	Email    string `gorm:"uniqueIndex;size:255;not null"`
	Password string `gorm:"size:255;not null"`
}
