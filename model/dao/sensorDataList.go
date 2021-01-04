package dao

import sensorDataDTO "github.com/naufalfmm/project-iot/model/dto/sensorData"

type SensorDataList []SensorData

func (sdl SensorDataList) Len() int {
	return len(sdl)
}

func NewFromCreatesDTO(ss []sensorDataDTO.CreateDTO) SensorDataList {
	ssLen := len(ss)
	sdlData := make(SensorDataList, len(ss))

	for i := 0; i < ssLen; i++ {
		sdlData[i] = NewFromCreateDTO(ss[i])
	}

	return sdlData
}

func (sdl SensorDataList) ToResponsesDTO() []sensorDataDTO.ResponseDTO {
	respDTO := make([]sensorDataDTO.ResponseDTO, sdl.Len())

	for i := 0; i < sdl.Len(); i++ {
		respDTO[i] = sdl[i].ToResponseDTO()
	}

	return respDTO
}
