package entity

import (
	"encoding/json"
	"time"
)

type Submission struct {
	ID           int             `json:"id" gorm:"primaryKey"`
	UserId       int             `json:"user_id"`
	Answer       json.RawMessage `json:"answer" gorm:"type:jsonb"`
	RiskScore    int             `json:"risk_score"`
	RiskCategory string          `json:"risk_category"`
	CreatedAt    time.Time       `json:"created_at"`
	UpdatedAt    time.Time       `json:"updated_at"`
}
