package dao

import (
	"fmt"
	"time"

	"github.com/naufalfmm/project-iot/common/consts"

	"gorm.io/gorm"

	nodeSensorDTO "github.com/naufalfmm/project-iot/model/dto/nodeSensor"
)

type NodeSensor struct {
	ID          uint64 `gorm:"PRIMARY_KEY"`
	Code        string `gorm:"not null"`
	Description string
	Category    string `gorm:"not null"`
	Unit        string `gorm:"not null"`
	GroupLabel  string `gorm:"not null"`
	GroupTh     uint32 `gorm:"not null"`
	NodeID      uint64 `gorm:"not null"`
	BaseModelSoftDeleted
}

func (NodeSensor) TableName() string {
	return "node_sensors"
}

func (ns *NodeSensor) BeforeCreate(tx *gorm.DB) error {
	if ns.Code == "" {
		ns.Code = fmt.Sprintf("%d-%d-%s", ns.NodeID, ns.GroupTh, ns.Category)
	}

	if ns.Unit == "" {
		ns.Unit = consts.UnitSensorStandard[ns.Category]
	}

	return nil
}

func NewSensorFromCreateDTO(c nodeSensorDTO.CreateDTO, groupTh uint32) NodeSensor {
	now := time.Now()

	return NodeSensor{
		Code:        c.Code,
		Description: c.Description,
		Category:    c.Category,
		Unit:        c.Unit,
		GroupLabel:  c.GroupLabel,
		GroupTh:     groupTh,
		NodeID:      c.NodeID,
		BaseModelSoftDeleted: BaseModelSoftDeleted{
			CreatedAt: now,
			CreatedBy: c.By,
		},
	}
}

func (ns NodeSensor) ToResponseDTO() nodeSensorDTO.ResponseDTO {
	return nodeSensorDTO.ResponseDTO{
		ID:          ns.ID,
		Code:        ns.Code,
		Description: ns.Description,
		Category:    ns.Category,
		Unit:        ns.Unit,
		GroupLabel:  ns.GroupLabel,
		GroupTh:     ns.GroupTh,
		NodeID:      ns.NodeID,
		CreatedAt:   ns.BaseModelSoftDeleted.CreatedAt,
		CreatedBy:   ns.BaseModelSoftDeleted.CreatedBy,
	}
}
