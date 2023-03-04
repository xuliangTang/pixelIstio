package vs

import "istio.io/client-go/pkg/apis/networking/v1alpha3"

// VsService @Service
type VsService struct {
	VsMap *VsMapStruct `inject:"-"`
}

func NewVsService() *VsService {
	return &VsService{}
}

func (this *VsService) ListByNs(ns string) []*v1alpha3.VirtualService {
	return this.VsMap.ListAll(ns)
}
