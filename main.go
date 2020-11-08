package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/naufalfmm/project-iot/domain"
	httpRest "github.com/naufalfmm/project-iot/http"
	"github.com/naufalfmm/project-iot/resource"
	"github.com/naufalfmm/project-iot/resource/config"
	"github.com/naufalfmm/project-iot/resource/jwt"
	sql "github.com/naufalfmm/project-iot/resource/sql"

	sensorDataRepo "github.com/naufalfmm/project-iot/domain/sensorData/repository"
	sensorDataServ "github.com/naufalfmm/project-iot/domain/sensorData/service"
	sensorDataHandler "github.com/naufalfmm/project-iot/handlers/sensorData"
	sensorDataCont "github.com/naufalfmm/project-iot/http/sensorData"
)

func main() {
	conf, err := config.NewConfig()
	if err != nil {
		panic(err)
	}

	db, err := sql.New(conf)
	if err != nil {
		panic(err)
	}

	jwt, err := jwt.New(conf)
	if err != nil {
		panic(err)
	}

	resource := resource.New(conf, db, jwt)

	sensorDataRepoNew, _ := sensorDataRepo.New(resource)
	sensorDataServNew, _ := sensorDataServ.New(resource, sensorDataRepoNew)

	domain := domain.Domain{
		SensorData: sensorDataServNew,
	}

	sensorDataHandNew, _ := sensorDataHandler.NewHandler(domain, resource)
	sensorDataController := sensorDataCont.Controller{
		SensorData: sensorDataHandNew,
		Resource:   resource,
	}

	controllers := httpRest.Controllers{
		SensorData: sensorDataController,
	}

	routes := httpRest.Routes{
		Controllers: controllers,
	}

	e := echo.New()

	routes.Register(e)

	e.Start(fmt.Sprintf(":%d", conf.ServerPort))

}
