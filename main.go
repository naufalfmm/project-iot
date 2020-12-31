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
	"github.com/naufalfmm/project-iot/resource/validator"
	vv9 "gopkg.in/go-playground/validator.v9"

	groupSensorRepo "github.com/naufalfmm/project-iot/domain/nodeSensor/subdomain/groupSensor/repository"
	groupSensorServ "github.com/naufalfmm/project-iot/domain/nodeSensor/subdomain/groupSensor/service"

	nodeSensorRepo "github.com/naufalfmm/project-iot/domain/nodeSensor/repository"
	nodeSensorServ "github.com/naufalfmm/project-iot/domain/nodeSensor/service"
	nodeSensorSubdomain "github.com/naufalfmm/project-iot/domain/nodeSensor/subdomain"
	nodeSensorHandler "github.com/naufalfmm/project-iot/handlers/nodeSensor"
	nodeSensorCont "github.com/naufalfmm/project-iot/http/nodeSensor"

	sensorDataRepo "github.com/naufalfmm/project-iot/domain/sensorData/repository"
	sensorDataServ "github.com/naufalfmm/project-iot/domain/sensorData/service"
	sensorDataHandler "github.com/naufalfmm/project-iot/handlers/sensorData"
	sensorDataCont "github.com/naufalfmm/project-iot/http/sensorData"

	nodeRepo "github.com/naufalfmm/project-iot/domain/node/repository"
	nodeServ "github.com/naufalfmm/project-iot/domain/node/service"
	nodeHandler "github.com/naufalfmm/project-iot/handlers/node"
	nodeCont "github.com/naufalfmm/project-iot/http/node"

	userRepo "github.com/naufalfmm/project-iot/domain/user/repository"
	userServ "github.com/naufalfmm/project-iot/domain/user/service"
	userHandler "github.com/naufalfmm/project-iot/handlers/user"
	userCont "github.com/naufalfmm/project-iot/http/user"

	authHandler "github.com/naufalfmm/project-iot/handlers/auth"
	authCont "github.com/naufalfmm/project-iot/http/auth"
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

	validatorGo := vv9.New()

	validator, err := validator.New(validatorGo)
	if err != nil {
		panic(err)
	}

	resource := resource.New(conf, db, jwt, &validator)

	groupSensorRepoNew, _ := groupSensorRepo.New(resource)
	groupSensorServNew, _ := groupSensorServ.New(resource, groupSensorRepoNew)

	nodeSensorSubdomainNew := nodeSensorSubdomain.Subdomain{
		GroupSensor: groupSensorServNew,
	}

	nodeSensorRepoNew, _ := nodeSensorRepo.New(resource)
	nodeSensorServNew, _ := nodeSensorServ.New(resource, nodeSensorRepoNew, nodeSensorSubdomainNew)

	sensorDataRepoNew, _ := sensorDataRepo.New(resource)
	sensorDataServNew, _ := sensorDataServ.New(resource, sensorDataRepoNew)

	nodeRepoNew, _ := nodeRepo.New(resource)
	nodeServNew, _ := nodeServ.New(resource, nodeRepoNew)

	userRepoNew, _ := userRepo.New(resource)
	userServNew, _ := userServ.New(resource, userRepoNew)

	domain := domain.Domain{
		Node:       nodeServNew,
		SensorData: sensorDataServNew,
		NodeSensor: nodeSensorServNew,
		User:       userServNew,
	}

	sensorDataHandNew, _ := sensorDataHandler.NewHandler(domain, resource)
	sensorDataController := sensorDataCont.Controller{
		SensorData: sensorDataHandNew,
		Resource:   resource,
	}

	nodeSensorHandNew, _ := nodeSensorHandler.NewHandler(domain, resource)
	nodeSensorController := nodeSensorCont.Controller{
		NodeSensor: nodeSensorHandNew,
		Resource:   resource,
	}

	nodeHandNew, _ := nodeHandler.NewHandler(domain, resource)
	nodeController := nodeCont.Controller{
		Node:     nodeHandNew,
		Resource: resource,
	}

	authHandNew, _ := authHandler.NewHandler(domain, resource)
	authController := authCont.Controller{
		Auth:     authHandNew,
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
		NodeSensor: nodeSensorController,
		Auth:       authController,
	}

	routes := httpRest.Routes{
		Controllers: controllers,
		Resource:    resource,
	}

	e := echo.New()
	e.Validator = resource.Validator

	routes.Register(e)

	e.Start(fmt.Sprintf(":%d", conf.ServerPort))

}
