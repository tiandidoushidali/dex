package svc

import (
	"github.com/tiandidoushidali/dex-dodo/backend/app/api/internal/config"
)

type ServiceContext struct {
	Config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
