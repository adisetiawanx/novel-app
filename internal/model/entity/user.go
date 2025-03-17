package entity

import (
	"database/sql"
	"github.com/google/uuid"
	"time"
)

type UserRole string // Define a custom type

const (
	Admin   UserRole = "admin"
	Visitor UserRole = "visitor"
)

type User struct {
	ID        uuid.UUID      `gorm:"type:uuid;primaryKey"`
	Name      string         `gorm:"size:255;not null"`
	Email     string         `gorm:"size:255;unique;not null"`
	Password  string         `gorm:"size:255;not null"`
	Phone     string         `gorm:"size:255;not null"`
	Profile   sql.NullString `gorm:"size:255"`
	Role      UserRole       `gorm:"type:varchar(10);not null;default:'visitor'"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`

	Tokens []Token `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
}

func (u *User) TableName() string {
	return "users"
}
