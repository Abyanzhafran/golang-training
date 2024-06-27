package entity

import (
	"encoding/json"
	"time"
)

type Submission struct {
	SubmissionID int64           `json:"id" gorm:"primaryKey"`
	UserId       int64           `json:"user_id`
	Answer       json.RawMessage `json:"answer" gorm:"type:jsonb"`
	RiskScore    int64           `json:"risk_score"`
	RiskCategory string          `json:"risk_category"`
	CreatedAt    time.Time       `json:"created_at"`
	UpdatedAt    time.Time       `json:"updated_at"`
}
