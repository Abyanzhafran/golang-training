package entity

import "time"

type User struct {
	ID         int64        `json:"id" gorm:"primaryKey"`
	Name       string       `json:"name"`
	Email      string       `json:"email"`
	CreatedAt  time.Time    `json:"created_at"`
	UpdatedAt  time.Time    `json:"updated_at"`
	Submission []Submission `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;`
}
