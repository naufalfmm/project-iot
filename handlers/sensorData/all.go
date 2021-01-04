package sensorData

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/naufalfmm/project-iot/common/consts"
	sensorDataDTO "github.com/naufalfmm/project-iot/model/dto/sensorData"
)

func (h *handler) All(ctx echo.Context, params sensorDataDTO.AllRequestParamsDTO) (sensorDataDTO.GetAllResponseDTO, error) {
	next, allRes, err := h.domain.SensorData.AllNext(ctx, params)
	if err != nil {
		ctx.Set(consts.ResponseCode, http.StatusInternalServerError)
		return sensorDataDTO.GetAllResponseDTO{}, err
	}

	allDTO := allRes.ToResponsesDTO()

	resp := sensorDataDTO.NewGetAllResponseDTO(params, allDTO, next)

	return resp, nil
}
