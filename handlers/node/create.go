package node

import (
	"github.com/labstack/echo/v4"
	"github.com/naufalfmm/project-iot/common/consts"

	nodeDTO "github.com/naufalfmm/project-iot/model/dto/node"
	sensorGroupDTO "github.com/naufalfmm/project-iot/model/dto/sensorGroup"
)

func (h *handler) Create(ctx echo.Context, createReq nodeDTO.CreateRequestDTO) (nodeDTO.CreateResponseDTO, error) {
	groupNumberLabel := createReq.Body.SensorGroupLabels
	groupNumber := len(groupNumberLabel)

	nodeCreateDTO := nodeDTO.CreateDTO{
		Label:       createReq.Body.Label,
		Location:    createReq.Body.Location,
		Type:        createReq.Body.Type,
		GroupNumber: uint64(groupNumber),
		By:          createReq.By.ID,
	}

	tx := h.resource.DB.Begin()

	ctx.Set(consts.PostgreTrx, tx)

	newNodeDTO, err := h.domain.Node.Create(ctx, nodeCreateDTO)
	if err != nil {
		tx.Rollback()
		return nodeDTO.CreateResponseDTO{}, err
	}

	sensorGroups := make([]sensorGroupDTO.ResponseDTO, groupNumber)

	for i := 0; i < groupNumber; i++ {
		groupCreateDTO := sensorGroupDTO.CreateDTO{
			Label:  groupNumberLabel[i],
			Th:     uint64(i),
			NodeID: newNodeDTO.ID,
		}

		newGroupDTO, err := h.domain.SensorGroup.Create(ctx, groupCreateDTO)
		if err != nil {
			tx.Rollback()
			return nodeDTO.CreateResponseDTO{}, err
		}

		sensorGroups[i] = newGroupDTO
	}

	resp := nodeDTO.CreateResponseDTO{
		ResponseDTO:  newNodeDTO,
		SensorGroups: sensorGroups,
	}

	tx.Commit()

	return resp, nil
}
