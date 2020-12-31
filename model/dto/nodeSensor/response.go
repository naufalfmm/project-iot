package nodeSensor

import "time"

type ResponseDTO struct {
	ID          uint64    `json:"id"`
	Code        string    `json:"code"`
	Description string    `json:"description"`
	Category    string    `json:"category"`
	Unit        string    `json:"unit"`
	GroupLabel  string    `json:"group_label"`
	GroupTh     uint32    `json:"group_th"`
	NodeID      uint64    `json:"node_id"`
	IsActive    bool      `json:"is_active"`
	CreatedAt   time.Time `json:"created_at"`
	CreatedBy   string    `json:"created_by"`
}
