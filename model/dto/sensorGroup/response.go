package sensorGroup

type ResponseDTO struct {
	ID     uint64 `json:"id"`
	Label  string `json:"label"`
	NodeID uint64 `json:"node_id"`
	Th     uint64 `json:"th"`
}
