package dao

type SensorGroup struct {
	BaseModelSoftDeleted
	ID     uint64 `gorm:"PRIMARY_KEY"`
	Label  string `gorm:"not null"`
	NodeID uint64 `gorm:"not null"`
	Th     uint64 `gorm:"not null"`
	Node   Node
}

func (SensorGroup) TableName() string {
	return "sensor_groups"
}

// func (sg SensorGroup) ToResponseDTO() sensorGroupDTO.ResponseDTO {
// 	return sensorGroupDTO.ResponseDTO(sg)
// }

// func NewSensorGroupFromCreateDTO(c sensorGroupDTO.CreateDTO) SensorGroup {
// 	now := time.Now()

// 	return SensorGroup{
// 		Label:     c.Label,
// 		NodeID:    c.NodeID,
// 		Th:        c.Th,
// 		CreatedAt: now,
// 	}
// }
