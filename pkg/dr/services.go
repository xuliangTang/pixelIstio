package dr

import (
	"istio.io/client-go/pkg/apis/networking/v1alpha3"
	"istio.io/client-go/pkg/clientset/versioned"
)

type DestinationRuleService struct {
	IstioClient        *versioned.Clientset `inject:"-"`
	DestinationRuleMap *DestinationRuleMap  `inject:"-"`
}

func NewDestinationRuleService() *DestinationRuleService {
	return &DestinationRuleService{}
}

func (this *DestinationRuleService) ListByNs(ns string) []*v1alpha3.DestinationRule {
	return this.DestinationRuleMap.ListByNs(ns)
}
