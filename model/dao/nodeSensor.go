package dao

import (
	"fmt"

	"github.com/naufalfmm/project-iot/common/word"
	"gorm.io/gorm"
)

type NodeSensor struct {
	ID          uint64 `gorm:"PRIMARY_KEY"`
	Code        string `gorm:"not null"`
	Description string `gorm:"not null"`
	Category    string `gorm:"not null"`
	Unit        string `gorm:"not null"`
	GroupLabel  string `gorm:"not null"`
	GroupTh     uint32 `gorm:"not null"`
	NodeID      uint64 `gorm:"not null"`
	BaseModelSoftDeleted
}

func (ns *NodeSensor) BeforeCreate(tx *gorm.DB) error {
	if ns.Code == "" {
		ns.Code = fmt.Sprintf("%d-%s-%s", ns.NodeID, word.FirstLetterWord(ns.GroupLabel), ns.Category)
	}

	return nil
}
