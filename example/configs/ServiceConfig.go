package configs

import "github.com/xingyunyang01/ioc/example/services"

type ServiceConfig struct {
}

func NewServiceConfig() *ServiceConfig {
	return &ServiceConfig{}
}

func (this *ServiceConfig) OrderService() *services.OrderService {
	return services.NewOrderService()
}
