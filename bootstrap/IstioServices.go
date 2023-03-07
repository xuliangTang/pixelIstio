package bootstrap

import (
	"pixelIstio/pkg/dr"
	"pixelIstio/pkg/gw"
	"pixelIstio/pkg/vs"
)

type IstioServiceConfig struct{}

func NewIstioServiceConfig() *IstioServiceConfig {
	return &IstioServiceConfig{}
}

func (*IstioServiceConfig) VsService() *vs.VsService {
	return vs.NewVsService()
}

func (*IstioServiceConfig) GwService() *gw.GateWayService {
	return gw.NewGateWayService()
}

func (*IstioServiceConfig) DrService() *dr.DestinationRuleService {
	return dr.NewDestinationRuleService()
}
