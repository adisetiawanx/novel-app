package domain

import "time"

type User struct {
	ID        uint      `gorm:"primary_key;auto_increment"`
	Name      string    `gorm:"size:255"`
	Email     string    `gorm:"size:255"`
	Phone     string    `gorm:"size:255"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoCreateTime;autoUpdateTime"`
}

func (u *User) TableName() string {
	return "users"
}
