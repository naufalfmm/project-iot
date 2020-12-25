package dao

type SensorGroupTypes []SensorGroupType

func (sgts SensorGroupTypes) Len() int {
	return len(sgts)
}
