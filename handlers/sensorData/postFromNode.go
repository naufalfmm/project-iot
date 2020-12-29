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

	sensorGroupTypeData, err := h.domain.SensorGroupType.AllByNodeID(e, nodeData.ID)
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
			GroupTh:    sensorGroupTypeData[i].SensorGroup.Th,
			SensorCode: sensorGroupTypeData[i].Code,
			SensorType: sensorGroupTypeData[i].Type,
			Value:      data[i],
			Unit:       sensorGroupTypeData[i].Unit,
			Timestamp:  req.Timestamp,
			CreatedBy:  nodeData.Label,
		}

		sensorDataReqs[i] = sensorDataReq

		resp.Data[i] = sensorDataReq.ToResponseDTO()
	}

	err = h.domain.SensorData.BulkInsert(e, sensorDataReqs)
	if err != nil {
		return sensorDataDTO.PostFromNodeResponseDTO{}, err
	}

	return resp, nil
}
