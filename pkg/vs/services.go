package vs

import (
	"context"
	"istio.io/client-go/pkg/apis/networking/v1alpha3"
	"istio.io/client-go/pkg/clientset/versioned"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// VsService @Service
type VsService struct {
	VsMap       *VsMapStruct         `inject:"-"`
	IstioClient *versioned.Clientset `inject:"-"`
}

func NewVsService() *VsService {
	return &VsService{}
}

func (this *VsService) ListByNs(ns string) []*v1alpha3.VirtualService {
	return this.VsMap.ListAll(ns)
}

func (this *VsService) Create(vs *v1alpha3.VirtualService) error {
	_, err := this.IstioClient.NetworkingV1alpha3().VirtualServices(vs.Namespace).Create(context.Background(), vs, v1.CreateOptions{})

	return err
}

func (this *VsService) Load(ns, name string) *v1alpha3.VirtualService {
	vs := this.VsMap.Get(ns, name)
	return vs
}

func (this *VsService) Update(vs *v1alpha3.VirtualService) error {
	oldVs := this.Load(vs.Namespace, vs.Name)
	vs.ResourceVersion = oldVs.ResourceVersion
	_, err := this.IstioClient.NetworkingV1alpha3().VirtualServices(vs.Namespace).Update(context.Background(), vs, v1.UpdateOptions{})
	return err
}

func (this *VsService) Delete(ns, name string) error {
	return this.IstioClient.NetworkingV1alpha3().VirtualServices(ns).Delete(context.Background(), name, v1.DeleteOptions{})
}
