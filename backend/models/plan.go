package models

import (
	"time"

	"github.com/google/uuid"
)

type Plan struct {
	PlanID         uuid.UUID     `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"plan_id"`
	PaymentMethod  PaymentMethod `json:"plan_payment_method"`
	ClubID         []Club        `json:"plan_club_id"`
	Name           string        `json:"plan_name"`
	Prince_in_cent int           `json:"plan_price"`
	Description    string        `json:"plan_description"`
	Active         bool          `json:"plan_status"`
	CreatedAt      time.Time     `json:"created_at"`
}
