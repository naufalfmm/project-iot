package dao

type SensorGroupType struct {
	ID            uint64 `gorm:"PRIMARY_KEY"`
	Code          string `gorm:"not null"`
	NodeID        uint64 `gorm:"not null"`
	SensorGroupID uint64 `gorm:"not null"`
	Description   string
	Type          string `gorm:"not null"`
	Unit          string `gorm:"not null"`
	BaseModelSoftDeleted
}

func (SensorGroupType) TableName() string {
	return "sensor_group_types"
}
