package sensorData

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/naufalfmm/project-iot/common/consts"

	sensorDataDTO "github.com/naufalfmm/project-iot/model/dto/sensorData"
)

func (h *handler) PostFromNode(ctx echo.Context, req sensorDataDTO.PostFromNodeRequestDTO) (sensorDataDTO.PostFromNodeResponseDTO, error) {
	var sensorDataReq sensorDataDTO.CreateDTO

	nodeData, err := h.domain.Node.CheckToken(ctx, req.Token)
	if err != nil {
		ctx.Set(consts.ResponseCode, http.StatusInternalServerError)
		return sensorDataDTO.PostFromNodeResponseDTO{}, err
	}

	nodeSensors, err := h.domain.NodeSensor.AllByNodeID(ctx, nodeData.ID)
	if err != nil {
		ctx.Set(consts.ResponseCode, http.StatusInternalServerError)
		return sensorDataDTO.PostFromNodeResponseDTO{}, err
	}

	if len(nodeSensors) < 1 {
		ctx.Set(consts.ResponseCode, http.StatusBadRequest)
		return sensorDataDTO.PostFromNodeResponseDTO{}, errors.New("No sensor assigned")
	}

	resp := sensorDataDTO.PostFromNodeResponseDTO{
		Node: nodeData.ToResponseDTO(),
		Data: make([]sensorDataDTO.ResponseDTO, len(nodeSensors)),
	}

	data := req.Data

	sensorDataReqs := make([]sensorDataDTO.CreateDTO, len(nodeSensors))

	for i := 0; i < len(nodeSensors); i++ {
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

	_, err = h.domain.SensorData.BulkInsert(ctx, sensorDataReqs)
	if err != nil {
		ctx.Set(consts.ResponseCode, http.StatusInternalServerError)
		return sensorDataDTO.PostFromNodeResponseDTO{}, err
	}

	return resp, nil
}
