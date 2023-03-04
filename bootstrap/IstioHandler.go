package bootstrap

import (
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
