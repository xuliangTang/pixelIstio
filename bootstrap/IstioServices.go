package bootstrap

import "pixelIstio/pkg/vs"

type IstioServiceConfig struct{}

func NewIstioServiceConfig() *IstioServiceConfig {
	return &IstioServiceConfig{}
}

func (*IstioServiceConfig) VsService() *vs.VsService {
	return vs.NewVsService()
}
