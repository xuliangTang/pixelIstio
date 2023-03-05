package bootstrap

import (
	"pixelIstio/pkg/gw"
	"pixelIstio/pkg/vs"
)

type IstioMaps struct {
}

func NewIstioMaps() *IstioMaps {
	return &IstioMaps{}
}

// InitVsMap 初始化 VsMapStruct
func (this *IstioMaps) InitVsMap() *vs.VsMapStruct {
	return &vs.VsMapStruct{}
}

func (this *IstioMaps) InitGwMap() *gw.GatewayMap {
	return &gw.GatewayMap{}
}
