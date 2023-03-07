package bootstrap

import (
	"pixelIstio/pkg/dr"
	"pixelIstio/pkg/gw"
	"pixelIstio/pkg/vs"
)

// IstioHandler 注入回调handler
type IstioHandler struct{}

func NewIstioHandler() *IstioHandler {
	return &IstioHandler{}
}

// VsHandler handler
func (this *IstioHandler) VsHandler() *vs.VsHandler {
	return &vs.VsHandler{}
}

func (this *IstioHandler) GwHandler() *gw.GateWayHandler {
	return &gw.GateWayHandler{}
}

func (this *IstioHandler) DrHandler() *dr.DestinationRuleHandler {
	return &dr.DestinationRuleHandler{}
}
