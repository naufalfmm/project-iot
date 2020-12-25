package dao

import "time"

type (
	BaseModel struct {
		CreatedAt time.Time `gorm:"not null"`
		CreatedBy string    `gorm:"not null"`
		UpdatedAt *time.Time
		UpdatedBy *string
	}

	BaseModelSoftDeleted struct {
		CreatedAt time.Time `gorm:"not null"`
		CreatedBy string    `gorm:"not null"`
		UpdatedAt *time.Time
		UpdatedBy *string
		DeletedAt *time.Time
		DeletedBy *string
		IsDeleted bool
	}
)
