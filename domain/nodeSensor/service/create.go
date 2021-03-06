package service

import (
	"github.com/jackc/pgconn"
	"github.com/labstack/echo/v4"
	"github.com/naufalfmm/project-iot/common/consts"
	"github.com/naufalfmm/project-iot/model/dao"
	nodeSensorDTO "github.com/naufalfmm/project-iot/model/dto/nodeSensor"
)

func (s *service) Create(ctx echo.Context, create nodeSensorDTO.CreateDTO) (dao.NodeSensor, error) {
	var groupTh uint32

	groups, err := s.subdomain.GroupSensor.GetUniqueByNodeID(ctx, create.NodeID)
	if err != nil {
		return dao.NodeSensor{}, err
	}

	groupTh = s.getGroupTh(groups, create.GroupLabel)

	newSensor := dao.NewSensorFromCreateDTO(create, groupTh)

	newSensor, err = s.repository.Create(ctx, newSensor)
	if err != nil {
		switch err.(type) {
		case *pgconn.PgError:
			{
				pqErr := err.(*pgconn.PgError)
				if pqErr.Code == "23505" {
					return dao.NodeSensor{}, consts.UniqueError
				}

				return dao.NodeSensor{}, err
			}
		default:
			{
				return dao.NodeSensor{}, err
			}
		}
	}

	return newSensor, nil
}

func (s *service) getGroupTh(groupsList dao.GroupSensors, groupLabel string) uint32 {
	for _, gr := range groupsList {
		if gr.GroupLabel == groupLabel {
			return gr.GroupTh
		}
	}

	return uint32(len(groupsList))
}
