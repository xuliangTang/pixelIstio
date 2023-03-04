package main

import (
	"github.com/xuliangTang/athena/athena"
	"pixelIstio/bootstrap"
	"pixelIstio/pkg/namespace"
	"pixelIstio/pkg/vs"
	"pixelIstio/pkg/wscore"
)

func main() {
	server := athena.Ignite().
		Configuration(
			bootstrap.NewK8sConfig(),
			bootstrap.NewIstioHandler(),
			bootstrap.NewIstioInformerStart(),
			bootstrap.NewIstioMaps(),
			bootstrap.NewIstioServiceConfig(),
		).Mount("", nil,
		vs.NewVsCtl(),
		namespace.NewNsCtl(),
		wscore.NewWsCtl(),
	)

	server.Launch()
}
