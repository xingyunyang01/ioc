package services

import "fmt"

type UserService struct {
	order *OrderService `inject:"-"`
}

func NewUserService() *UserService {
	return &UserService{}
}

func (this *UserService) GetUserInfo(uid int) {
	fmt.Println("获取用户id=", uid, "的详细信息")
}

func (this *UserService) GetOrderInfo(uid int) {
	this.order.GetOrderInfo(uid)
}
