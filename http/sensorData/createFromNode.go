package sensorData

import (
	"errors"
	"net/http"
	"regexp"
	"strconv"

	"github.com/naufalfmm/project-iot/common/defaultResp"

	"github.com/naufalfmm/project-iot/common/consts"

	"github.com/labstack/echo/v4"
	sensorDataDTO "github.com/naufalfmm/project-iot/model/dto/sensorData"
)

func (c *Controller) CreateFromNode(e echo.Context) error {
	req, err := c.paramBind(e)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	err = e.Validate(req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	resp, err := c.SensorData.PostFromNode(e, req)
	if err != nil {
		if errors.Is(err, consts.Unauthorized) {
			return echo.NewHTTPError(http.StatusUnauthorized, err)
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return e.JSON(http.StatusCreated, defaultResp.CreateSuccessResp(http.StatusCreated, resp))
}

func (c *Controller) paramBind(e echo.Context) (sensorDataDTO.PostFromNodeRequestDTO, error) {
	qp := e.QueryParams()
	req := sensorDataDTO.PostFromNodeRequestDTO{}

	for k := range qp {
		if k == "token" {
			req.Token = qp.Get(k)
		}

		if !c.isValidParamKey(k) {
			return sensorDataDTO.PostFromNodeRequestDTO{}, errors.New("Wrong param key format")
		}

		key, i, err := c.splitParamKeyWithIndex(k)
		if err != nil {
			return sensorDataDTO.PostFromNodeRequestDTO{}, err
		}

		req = c.buildPostDTO(req, i)

		data, err := strconv.ParseFloat(qp.Get(k), 32)
		if err != nil {
			return sensorDataDTO.PostFromNodeRequestDTO{}, err
		}

		if key == "temp" {
			req.Data[i].Temp = data
			continue
		}

		if key == "tds" {
			req.Data[i].TDS = data
			continue
		}

		if key == "ph" {
			req.Data[i].PH = data
			continue
		}
	}

	return req, nil
}

func (c *Controller) isValidParamKey(param string) bool {
	re := regexp.MustCompile(`^[a-zA-Z]+\d$`)

	return re.Match([]byte(param))
}

func (c *Controller) splitParamKeyWithIndex(paramKey string) (string, int, error) {
	re := regexp.MustCompile(`^([a-zA-Z]+)(\d)$`)

	res := re.FindStringSubmatch(paramKey)
	if len(res) < 3 {
		return "", 0, errors.New("Wrong param key format")
	}

	idx, err := strconv.Atoi(res[2])
	if err != nil {
		return "", 0, err
	}

	return res[1], idx, nil
}

func (c *Controller) buildPostDTO(req sensorDataDTO.PostFromNodeRequestDTO, idx int) sensorDataDTO.PostFromNodeRequestDTO {
	len := len(req.Data)

	for i := 0; i < idx-len+1; i++ {
		req.Data = append(req.Data, sensorDataDTO.PostDTO{})
	}

	return req
}
