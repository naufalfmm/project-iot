package node

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/naufalfmm/project-iot/common/consts"

	nodeDTO "github.com/naufalfmm/project-iot/model/dto/node"
)

func (h *handler) Create(ctx echo.Context, createReq nodeDTO.CreateRequestDTO) (nodeDTO.ResponseDTO, error) {
	nodeCreateDTO := nodeDTO.CreateDTO{
		Label:    createReq.Body.Label,
		Location: createReq.Body.Location,
		Type:     createReq.Body.Type,
		By:       createReq.By.Username,
	}

	newNode, err := h.domain.Node.Create(ctx, nodeCreateDTO)
	if err != nil {
		statusCode := http.StatusInternalServerError
		if errors.Is(err, consts.UniqueError) {
			statusCode = http.StatusConflict
		}

		ctx.Set(consts.ResponseCode, statusCode)
		return nodeDTO.ResponseDTO{}, err
	}

	return newNode.ToResponseDTO(), nil
}
