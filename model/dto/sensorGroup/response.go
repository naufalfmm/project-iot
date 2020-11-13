package sensorGroup

import "time"

type ResponseDTO struct {
	ID        uint64     `json:"id"`
	Label     string     `json:"label"`
	NodeID    uint64     `json:"node_id"`
	Th        uint64     `json:"th"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
	IsDeleted bool       `json:"is_deleted"`
}
