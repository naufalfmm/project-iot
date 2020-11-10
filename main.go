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

	sensorGroupRepo "github.com/naufalfmm/project-iot/domain/sensorGroup/repository"
	sensorGroupServ "github.com/naufalfmm/project-iot/domain/sensorGroup/service"

	nodeRepo "github.com/naufalfmm/project-iot/domain/node/repository"
	nodeServ "github.com/naufalfmm/project-iot/domain/node/service"
	nodeHandler "github.com/naufalfmm/project-iot/handlers/node"
	nodeCont "github.com/naufalfmm/project-iot/http/node"

	userRepo "github.com/naufalfmm/project-iot/domain/user/repository"
	userServ "github.com/naufalfmm/project-iot/domain/user/service"
	userHandler "github.com/naufalfmm/project-iot/handlers/user"
	userCont "github.com/naufalfmm/project-iot/http/user"
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

	sensorGroupRepoNew, _ := sensorGroupRepo.New(resource)
	sensorGroupServNew, _ := sensorGroupServ.New(resource, sensorGroupRepoNew)

	nodeRepoNew, _ := nodeRepo.New(resource)
	nodeServNew, _ := nodeServ.New(resource, nodeRepoNew)

	userRepoNew, _ := userRepo.New(resource)
	userServNew, _ := userServ.New(resource, userRepoNew)

	domain := domain.Domain{
		Node:        nodeServNew,
		SensorData:  sensorDataServNew,
		SensorGroup: sensorGroupServNew,
		User:        userServNew,
	}

	sensorDataHandNew, _ := sensorDataHandler.NewHandler(domain, resource)
	sensorDataController := sensorDataCont.Controller{
		SensorData: sensorDataHandNew,
		Resource:   resource,
	}

	nodeHandNew, _ := nodeHandler.NewHandler(domain, resource)
	nodeController := nodeCont.Controller{
		Node:     nodeHandNew,
		Resource: resource,
	}

	userHandNew, _ := userHandler.NewHandler(domain, resource)
	userController := userCont.Controller{
		User:     userHandNew,
		Resource: resource,
	}

	controllers := httpRest.Controllers{
		SensorData: sensorDataController,
		Node:       nodeController,
		User:       userController,
	}

	routes := httpRest.Routes{
		Controllers: controllers,
		Resource:    resource,
	}

	e := echo.New()

	routes.Register(e)

	e.Start(fmt.Sprintf(":%d", conf.ServerPort))

}
