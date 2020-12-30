package sensorData

import (
	"github.com/labstack/echo/v4"

	sensorDataDTO "github.com/naufalfmm/project-iot/model/dto/sensorData"
)

func (h *handler) PostFromNode(e echo.Context, req sensorDataDTO.PostFromNodeRequestDTO) (sensorDataDTO.PostFromNodeResponseDTO, error) {
	var sensorDataReq sensorDataDTO.CreateDTO

	nodeData, err := h.domain.Node.CheckToken(e, req.Token)
	if err != nil {
		return sensorDataDTO.PostFromNodeResponseDTO{}, err
	}

	nodeSensors, err := h.domain.NodeSensor.AllByNodeID(e, nodeData.ID)
	if err != nil {
		return sensorDataDTO.PostFromNodeResponseDTO{}, err
	}

	resp := sensorDataDTO.PostFromNodeResponseDTO{
		Node: nodeData.ToResponseDTO(),
		Data: make([]sensorDataDTO.ResponseDTO, len(req.Data)),
	}

	data := req.Data

	sensorDataReqs := make([]sensorDataDTO.CreateDTO, len(req.Data))

	for i := 0; i < len(data); i++ {
		sensorDataReq = sensorDataDTO.CreateDTO{
			NodeID:     nodeData.ID,
			NodeLabel:  nodeData.Label,
			Code:       nodeSensors[i].Code,
			Category:   nodeSensors[i].Category,
			Value:      data[i],
			Unit:       nodeSensors[i].Unit,
			GroupTh:    nodeSensors[i].GroupTh,
			GroupLabel: nodeSensors[i].GroupLabel,
			Timestamp:  req.Timestamp,
			CreatedBy:  nodeData.Label,
		}

		sensorDataReqs[i] = sensorDataReq

		resp.Data[i] = sensorDataReq.ToResponseDTO()
	}

	_, err = h.domain.SensorData.BulkInsert(e, sensorDataReqs)
	if err != nil {
		return sensorDataDTO.PostFromNodeResponseDTO{}, err
	}

	return resp, nil
}
