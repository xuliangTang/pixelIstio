package dr

import (
	"context"
	"istio.io/client-go/pkg/apis/networking/v1alpha3"
	"istio.io/client-go/pkg/clientset/versioned"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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

func (this *DestinationRuleService) Delete(ns, name string) error {
	return this.IstioClient.NetworkingV1alpha3().DestinationRules(ns).Delete(context.Background(), name, v1.DeleteOptions{})
}

func (this *DestinationRuleService) Show(ns, name string) *v1alpha3.DestinationRule {
	return this.DestinationRuleMap.Get(ns, name)
}

func (this *DestinationRuleService) Create(dr *v1alpha3.DestinationRule) error {
	_, err := this.IstioClient.NetworkingV1alpha3().DestinationRules(dr.Namespace).Create(context.Background(), dr, v1.CreateOptions{})
	return err
}

func (this *DestinationRuleService) Update(dr *v1alpha3.DestinationRule) error {
	oldDr := this.Show(dr.Namespace, dr.Name)
	dr.ResourceVersion = oldDr.ResourceVersion
	_, err := this.IstioClient.NetworkingV1alpha3().DestinationRules(dr.Namespace).Update(context.Background(), dr, v1.UpdateOptions{})
	return err
}
