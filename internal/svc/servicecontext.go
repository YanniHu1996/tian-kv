package svc

import (
	config2 "tian-kv/internal/config"
)

type ServiceContext struct {
	Config config2.Config
}

func NewServiceContext(c config2.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
