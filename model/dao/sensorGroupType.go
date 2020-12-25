package dao

type SensorGroupType struct {
	BaseModelSoftDeleted
	ID            uint64 `gorm:"PRIMARY_KEY"`
	Code          string `gorm:"not null"`
	SensorGroupID uint64 `gorm:"not null"`
	Description   string
	Type          string `gorm:"not null"`
	Unit          string `gorm:"not null"`
	SensorGroup   SensorGroup
}

func (SensorGroupType) TableName() string {
	return "sensor_group_types"
}
