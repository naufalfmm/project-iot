package dao

import (
	"time"

	"github.com/naufalfmm/project-iot/common/token"
	nodeDTO "github.com/naufalfmm/project-iot/model/dto/node"
)

type Node struct {
	ID          uint64 `gorm:"PRIMARY_KEY"`
	Label       string `gorm:"not null"`
	Location    *string
	Token       string    `gorm:"not null"`
	Type        string    `gorm:"not null"`
	GroupNumber uint64    `gorm:"not null"`
	CreatedAt   time.Time `gorm:"not null"`
	CreatedBy   uint64    `gorm:"not null"`
	UpdatedAt   *time.Time
	UpdatedBy   *uint64
	DeletedAt   *time.Time
	DeletedBy   *time.Time
	IsDeleted   bool
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
		IsDeleted: n.IsDeleted,
	}
}

func NewNodeFromCreateDTO(r nodeDTO.CreateDTO) Node {
	now := time.Now()

	return Node{
		Label:       r.Label,
		Location:    r.Location,
		Token:       token.GenerateNodeToken(r.Label),
		Type:        r.Type,
		GroupNumber: r.GroupNumber,
		CreatedAt:   now,
		CreatedBy:   r.By,
	}
}
