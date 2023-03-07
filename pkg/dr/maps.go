package dr

import (
	"fmt"
	"istio.io/client-go/pkg/apis/networking/v1alpha3"
	"sort"
	"sync"
)

type DestinationRuleMap struct {
	Data sync.Map // [ns string] []*DestinationRule
}

func (this *DestinationRuleMap) Get(ns, name string) *v1alpha3.DestinationRule {
	if list, ok := this.Data.Load(ns); ok {
		drList := list.([]*v1alpha3.DestinationRule)
		for _, dr := range drList {
			if dr.Name == name {
				return dr
			}
		}
	}
	return nil
}

func (this *DestinationRuleMap) Add(dr *v1alpha3.DestinationRule) {
	if list, ok := this.Data.Load(dr.Namespace); ok {
		drList := list.([]*v1alpha3.DestinationRule)
		newList := append(drList, dr)
		this.Data.Store(dr.Namespace, newList)
	} else {
		this.Data.Store(dr.Namespace, []*v1alpha3.DestinationRule{dr})
	}
}

func (this *DestinationRuleMap) Update(dr *v1alpha3.DestinationRule) error {
	if list, ok := this.Data.Load(dr.Namespace); ok {
		drList := list.([]*v1alpha3.DestinationRule)
		for i, item := range drList {
			if item.Name == dr.Name {
				drList[i] = dr
				return nil
			}
		}
	}
	return fmt.Errorf("DestinationRule-%s not found", dr.Name)
}

func (this *DestinationRuleMap) Delete(dr *v1alpha3.DestinationRule) {
	if list, ok := this.Data.Load(dr.Namespace); ok {
		drList := list.([]*v1alpha3.DestinationRule)
		for i, item := range drList {
			if item.Name == dr.Name {
				newDrList := append(drList[:i], drList[i+1:]...)
				this.Data.Store(dr.Namespace, newDrList)
				break
			}
		}
	}
}

func (this *DestinationRuleMap) ListByNs(ns string) []*v1alpha3.DestinationRule {
	if list, ok := this.Data.Load(ns); ok {
		drList := list.([]*v1alpha3.DestinationRule)
		sort.Sort(DR(drList))
		return drList
	}
	return []*v1alpha3.DestinationRule{}
}

type DR []*v1alpha3.DestinationRule

func (this DR) Len() int {
	return len(this)
}
func (this DR) Less(i, j int) bool {
	// 根据时间排序
	return this[i].CreationTimestamp.Time.After(this[j].CreationTimestamp.Time)
}
func (this DR) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}
