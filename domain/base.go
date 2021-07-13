package domain

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"time"
)

type Base struct {
	ID        string    `json:"id" gorm:"type:uuid;primary_key"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at" sql:"index"`
}

// BeforeCreate pertence ao gorm
func (base *Base) BeforeCreate(tx *gorm.DB) error {
	base.ID = uuid.NewV4().String()
	base.CreatedAt = time.Now()

	return nil
}
