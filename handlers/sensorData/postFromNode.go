package sensorData

import (
	"errors"
	"net/http"

	"github.com/naufalfmm/project-iot/model/dao"

	"github.com/labstack/echo/v4"
	"github.com/naufalfmm/project-iot/common/consts"

	sensorDataDTO "github.com/naufalfmm/project-iot/model/dto/sensorData"
)

func (h *handler) PostFromNode(ctx echo.Context, req sensorDataDTO.PostFromNodeRequestDTO) (sensorDataDTO.PostFromNodeResponseDTO, error) {
	var (
		sensorDataReq sensorDataDTO.CreateDTO
		nodeSensor    dao.NodeSensor
	)

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
	}

	data := req.Data

	dataResp := []sensorDataDTO.ResponseDTO{}
	sensorDataReqs := []sensorDataDTO.CreateDTO{}

	for i := 0; i < len(nodeSensors); i++ {
		nodeSensor = nodeSensors[i]

		if !nodeSensor.IsActive {
			continue
		}

		sensorDataReq = sensorDataDTO.CreateDTO{
			NodeID:     nodeData.ID,
			NodeLabel:  nodeData.Label,
			Code:       nodeSensor.Code,
			Category:   nodeSensor.Category,
			Value:      data[i],
			Unit:       nodeSensor.Unit,
			GroupTh:    nodeSensor.GroupTh,
			GroupLabel: nodeSensor.GroupLabel,
			Timestamp:  req.Timestamp,
			CreatedBy:  nodeData.Label,
		}

		sensorDataReqs = append(sensorDataReqs, sensorDataReq)

		dataResp = append(dataResp, sensorDataReq.ToResponseDTO())
	}

	_, err = h.domain.SensorData.BulkInsert(ctx, sensorDataReqs)
	if err != nil {
		ctx.Set(consts.ResponseCode, http.StatusInternalServerError)
		return sensorDataDTO.PostFromNodeResponseDTO{}, err
	}

	resp.Data = dataResp

	return resp, nil
}
