package dao

type GroupSensor struct {
	GroupLabel string `gorm:"not null"`
	GroupTh    uint32 `gorm:"not null"`
	NodeID     uint64 `gorm:"not null"`
}

func (GroupSensor) TableName() string {
	return "node_sensors"
}
