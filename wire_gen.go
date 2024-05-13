// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"LinkMe/internal/dao"
	"LinkMe/internal/repository"
	"LinkMe/internal/service"
	"LinkMe/internal/web"
	"LinkMe/ioc"
	"github.com/gin-gonic/gin"
)

import (
	_ "github.com/google/wire"
)

// Injectors from wire.go:

func InitWebServer() *gin.Engine {
	db := ioc.InitDB()
	userDAO := dao.NewUserDAO(db)
	userRepository := repository.NewUserRepository(userDAO)
	userService := service.NewUserService(userRepository)
	userHandler := web.NewUserHandler(userService)
	engine := ioc.InitWebServer(userHandler)
	return engine
}
