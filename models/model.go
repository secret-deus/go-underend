package models

type User struct {
	ID    uint   `gorm:"primary_key"`
	Name  string `gorm:"not null"`
	Email string `gorm:"not null;unique"`
}

type Resource struct {
	ID    uint   `gorm:"primary_key"`
	Name  string `gorm:"not null"`
	Type  string `gorm:"not null"`
	State string `gorm:"not null"`
}
