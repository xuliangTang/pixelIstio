package main

import (
	"github.com/xuliangTang/athena/athena"
	"pixelIstio/bootstrap"
	"pixelIstio/pkg/vs"
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
	)

	server.Launch()
}
