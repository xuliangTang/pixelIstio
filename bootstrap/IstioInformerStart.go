package bootstrap

import (
	istio "istio.io/client-go/pkg/clientset/versioned"
	"istio.io/client-go/pkg/informers/externalversions"
	"k8s.io/apimachinery/pkg/util/wait"
	"pixelIstio/pkg/gw"
	"pixelIstio/pkg/vs"
)

type IstioInformerStart struct {
	IstioClient *istio.Clientset   `inject:"-"`
	VsHandler   *vs.VsHandler      `inject:"-"`
	GwHandler   *gw.GateWayHandler `inject:"-"`
}

func NewIstioInformerStart() *IstioInformerStart {
	return &IstioInformerStart{}
}

// InitInformer 初始化Informer
func (this *IstioInformerStart) InitInformer() externalversions.SharedInformerFactory {
	fact := externalversions.NewSharedInformerFactoryWithOptions(this.IstioClient, 0)

	// 虚拟服务的监听
	fact.Networking().V1alpha3().VirtualServices().Informer().AddEventHandler(this.VsHandler)
	// 网关的监听
	fact.Networking().V1alpha3().Gateways().Informer().AddEventHandler(this.GwHandler)

	fact.Start(wait.NeverStop)

	return fact
}
