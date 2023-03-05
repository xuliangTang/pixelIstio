package gw

import (
	"context"
	"istio.io/client-go/pkg/apis/networking/v1alpha3"
	"istio.io/client-go/pkg/clientset/versioned"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type GateWayService struct {
	IstioClient *versioned.Clientset `inject:"-"`
	GatewayMap  *GatewayMap          `inject:"-"`
}

func NewGateWayService() *GateWayService {
	return &GateWayService{}
}

func (this *GateWayService) ListByNs(ns string) []*v1alpha3.Gateway {
	return this.GatewayMap.ListByNs(ns)
}

func (this *GateWayService) Load(ns, name string) *v1alpha3.Gateway {
	gw := this.GatewayMap.Get(ns, name)
	return gw
}

func (this *GateWayService) Create(gateway *v1alpha3.Gateway) error {
	_, err := this.IstioClient.NetworkingV1alpha3().Gateways(gateway.Namespace).Create(context.Background(), gateway, v1.CreateOptions{})
	return err
}

func (this *GateWayService) Delete(ns, name string) error {
	return this.IstioClient.NetworkingV1alpha3().Gateways(ns).Delete(context.Background(), name, v1.DeleteOptions{})
}

func (this *GateWayService) Show(ns, name string) *v1alpha3.Gateway {
	return this.GatewayMap.Get(ns, name)
}

func (this *GateWayService) Update(gateway *v1alpha3.Gateway) error {
	oldGW := this.Load(gateway.Namespace, gateway.Name)
	gateway.ResourceVersion = oldGW.ResourceVersion
	_, err := this.IstioClient.NetworkingV1alpha3().Gateways(gateway.Namespace).Update(context.Background(), gateway, v1.UpdateOptions{})
	return err
}
