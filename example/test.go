package main

import (
	"github.com/xingyunyang01/ioc"
	"github.com/xingyunyang01/ioc/example/configs"
	"github.com/xingyunyang01/ioc/example/services"
)

func main() {
	serviceConfig := configs.NewServiceConfig()

	ioc.BeanFactory.Config(serviceConfig)

	userService := services.NewUserService()
	ioc.BeanFactory.Apply(userService)
	userService.GetOrderInfo(1)
}
