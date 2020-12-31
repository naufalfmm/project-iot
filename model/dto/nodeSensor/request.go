package nodeSensor

import "github.com/naufalfmm/project-iot/common/login"

type (
	CreateDTO struct {
		Code        string
		Description string
		Category    string
		Unit        string
		GroupLabel  string
		NodeID      uint64
		By          string
	}
)

type (
	CreateRequestBodyDTO struct {
		Code        string `json:"code" validate:"required"`
		Description string `json:"description"`
		Category    string `json:"category" validate:"required"`
		Unit        string `json:"unit"`
		GroupLabel  string `json:"group_label" validate:"required"`
		NodeID      uint64 `json:"node_id" validate:"required"`
	}
)

type (
	CreateRequestDTO struct {
		Body CreateRequestBodyDTO
		By   login.ClientJWTDTO
	}
)

func (cr CreateRequestDTO) ToCreateDTO() CreateDTO {
	return CreateDTO{
		Code:        cr.Body.Code,
		Description: cr.Body.Description,
		Category:    cr.Body.Category,
		Unit:        cr.Body.Unit,
		GroupLabel:  cr.Body.GroupLabel,
		NodeID:      cr.Body.NodeID,
		By:          cr.By.Username,
	}
}
