package models

import (
	"time"

	"github.com/google/uuid"
)

type PaymentMethod struct {
	MethodID    uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"payment_method_id"`
	ClubID      []Club
	Name        string    `json:"payment_name"`
	Description string    `json:"payment_description"`
	CreatedAt   time.Time `json:"created_at"`
}
