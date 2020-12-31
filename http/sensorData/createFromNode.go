package sensorData

import (
	"errors"
	"net/http"
	"regexp"
	"strconv"
	"time"

	"github.com/naufalfmm/project-iot/common/defaultResp"

	"github.com/naufalfmm/project-iot/common/consts"

	"github.com/labstack/echo/v4"
	sensorDataDTO "github.com/naufalfmm/project-iot/model/dto/sensorData"
)

func (c *Controller) CreateFromNode(ctx echo.Context) error {
	req, err := c.paramBind(ctx)
	if err != nil {
		ctx.Set(consts.ResponseCode, http.StatusBadRequest)
		return defaultResp.CreateResp(ctx, err.Error())
	}

	err = ctx.Validate(req)
	if err != nil {
		ctx.Set(consts.ResponseCode, http.StatusBadRequest)
		return defaultResp.CreateResp(ctx, err.Error())
	}

	resp, err := c.SensorData.PostFromNode(ctx, req)
	if err != nil {
		return defaultResp.CreateResp(ctx, err.Error())
	}

	ctx.Set(consts.ResponseCode, http.StatusCreated)
	return defaultResp.CreateResp(ctx, resp)
}

func (c *Controller) paramBind(e echo.Context) (sensorDataDTO.PostFromNodeRequestDTO, error) {
	qp := e.QueryParams()
	req := sensorDataDTO.PostFromNodeRequestDTO{}

	now := time.Now()

	sensorData := []float64{}

	for k := range qp {
		if k == "token" {
			req.Token = qp.Get(k)
			continue
		}

		if !c.isValidParamKey(k) {
			return sensorDataDTO.PostFromNodeRequestDTO{}, errors.New("Wrong param key format")
		}

		data, err := strconv.ParseFloat(qp.Get(k), 64)
		if err != nil {
			return sensorDataDTO.PostFromNodeRequestDTO{}, err
		}

		sensorData = append(sensorData, data)
	}

	req.Data = sensorData
	req.Timestamp = now

	return req, nil
}

func (c *Controller) isValidParamKey(param string) bool {
	re := regexp.MustCompile(`^sensor\d$`)

	return re.Match([]byte(param))
}
