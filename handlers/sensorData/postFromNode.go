package sensorData

import (
	"github.com/labstack/echo/v4"
	"github.com/naufalfmm/project-iot/common/consts"

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
		Node: nodeData,
		Data: make([]sensorDataDTO.ResponseDTO, len(req.Data)),
	}

	data := req.Data

	tx := h.resource.DB.Begin()
	e.Set(consts.PostgreTrx, tx)

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

		sensorData, err := h.domain.SensorData.Insert(e, sensorDataReq)
		if err != nil {
			tx.Rollback()
			return sensorDataDTO.PostFromNodeResponseDTO{}, err
		}

		resp.Data[i] = sensorData
	}

	tx.Commit()

	return resp, nil
}
