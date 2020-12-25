package dao

type SensorGroups []SensorGroup

func (sgs SensorGroups) Len() int {
	return len(sgs)
}

// func (sgs SensorGroups) ToResponsesDTO() []sensorGroupDTO.ResponseDTO {
// 	sgsLen := sgs.Len()

// 	sgsDTO := make([]sensorGroupDTO.ResponseDTO, sgsLen)

// 	for i := 0; i < sgsLen; i++ {
// 		sgsDTO[i] = sgs[i].ToResponseDTO()
// 	}

// 	return sgsDTO
// }
