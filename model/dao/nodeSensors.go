package dao

type NodeSensors []NodeSensor

func (NodeSensors) TableName() string {
	return "node_sensors"
}
