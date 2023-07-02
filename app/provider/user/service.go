package user

import (
	"github.com/gohade/hade/framework"
	"github.com/gohade/hade/framework/contract"
)

type UserService struct {
	container framework.Container
	logger    contract.Log
	configer  contract.Config
}

func NewUserService(params ...interface{}) (interface{}, error) {
	container := params[0].(framework.Container)
	logger := container.MustMake(contract.LogKey).(contract.Log)
	configer := container.MustMake(contract.ConfigKey).(contract.Config)
	return &UserService{container: container, logger: logger, configer: configer}, nil
}
