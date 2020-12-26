package dao

import (
	"time"

	"github.com/naufalfmm/project-iot/common/token"
	nodeDTO "github.com/naufalfmm/project-iot/model/dto/node"
)

type Node struct {
	ID          uint64 `gorm:"PRIMARY_KEY"`
	Type        string `gorm:"not null"`
	Label       string `gorm:"not null"`
	Token       string `gorm:"not null"`
	Location    *string
	GroupNumber uint64 `gorm:"not null"`
	BaseModelSoftDeleted
}

func (Node) TableName() string {
	return "nodes"
}

func (n Node) ToResponseDTO() nodeDTO.ResponseDTO {
	return nodeDTO.ResponseDTO{
		ID:          n.ID,
		Label:       n.Label,
		Location:    n.Location,
		Token:       n.Token,
		Type:        n.Type,
		GroupNumber: n.GroupNumber,
		CreatedAt:   n.CreatedAt,
		CreatedBy:   n.CreatedBy,
		UpdatedAt:   n.UpdatedAt,
		UpdatedBy:   n.UpdatedBy,
		DeletedAt:   n.DeletedAt,
		DeletedBy:   n.DeletedBy,
		IsDeleted:   n.IsDeleted,
	}
}

func NewNodeFromCreateDTO(r nodeDTO.CreateDTO) Node {
	now := time.Now()

	return Node{
		Type:        r.Type,
		Label:       r.Label,
		Location:    r.Location,
		Token:       token.GenerateNodeToken(r.Label),
		GroupNumber: r.GroupNumber,
		BaseModelSoftDeleted: BaseModelSoftDeleted{
			CreatedAt: now,
			CreatedBy: r.By,
		},
	}
}
