package sensorData

import (
	"github.com/labstack/echo/v4"
	"github.com/naufalfmm/project-iot/common/consts"

	sensorDataDTO "github.com/naufalfmm/project-iot/model/dto/sensorData"
)

func (h *handler) PostFromNode(e echo.Context, req sensorDataDTO.PostFromNodeRequestDTO) (sensorDataDTO.PostFromNodeResponseDTO, error) {
	nodeData, err := h.domain.Node.CheckToken(e, req.Token)
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
		sensorDataReq := data[i].ToCreateDTO(uint64(i), nodeData.ID)

		sensorData, err := h.domain.SensorData.Insert(e, sensorDataReq, nodeData.ID)
		if err != nil {
			tx.Rollback()
			return sensorDataDTO.PostFromNodeResponseDTO{}, err
		}

		resp.Data[i] = sensorData
	}

	tx.Commit()

	return resp, nil
}
