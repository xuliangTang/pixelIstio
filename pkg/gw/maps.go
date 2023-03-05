package gw

import (
	"fmt"
	"istio.io/client-go/pkg/apis/networking/v1alpha3"
	"sort"
	"sync"
)

type GatewayMap struct {
	Data sync.Map // [ns string] []*Gateway
}

func (this *GatewayMap) Get(ns, name string) *v1alpha3.Gateway {
	if list, ok := this.Data.Load(ns); ok {
		for _, item := range list.([]*v1alpha3.Gateway) {
			if item.Name == name {
				return item
			}
		}
	}
	return nil
}

func (this *GatewayMap) Add(gw *v1alpha3.Gateway) {
	if list, ok := this.Data.Load(gw.Namespace); ok {
		list := append(list.([]*v1alpha3.Gateway), gw)
		this.Data.Store(gw.Namespace, list)
	} else {
		this.Data.Store(gw.Namespace, []*v1alpha3.Gateway{gw})
	}
}

func (this *GatewayMap) Update(gw *v1alpha3.Gateway) error {
	if list, ok := this.Data.Load(gw.Namespace); ok {
		gwList := list.([]*v1alpha3.Gateway)
		for i, item := range gwList {
			if item.Name == gw.Name {
				gwList[i] = gw
			}
		}
		return nil
	}
	return fmt.Errorf("gateway-%s not found", gw.Name)
}

func (this *GatewayMap) Delete(gw *v1alpha3.Gateway) {
	if list, ok := this.Data.Load(gw.Namespace); ok {
		gwList := list.([]*v1alpha3.Gateway)
		for i, item := range gwList {
			if item.Name == gw.Name {
				newGwList := append(gwList[:i], gwList[i+1:]...)
				this.Data.Store(gw.Namespace, newGwList)
				break
			}
		}
	}
}

func (this *GatewayMap) ListByNs(ns string) []*v1alpha3.Gateway {
	if list, ok := this.Data.Load(ns); ok {
		gwList := list.([]*v1alpha3.Gateway)
		sort.Sort(GW(gwList))
		return gwList
	}
	return []*v1alpha3.Gateway{}
}

func (this *GatewayMap) ListAll() []map[string]interface{} {
	ret := make([]map[string]interface{}, 0)

	this.Data.Range(func(key, value interface{}) bool {
		m := map[string]interface{}{
			"ns":   key,
			"list": value,
		}
		ret = append(ret, m)
		return true
	})

	return ret
}

type GW []*v1alpha3.Gateway

func (this GW) Len() int {
	return len(this)
}
func (this GW) Less(i, j int) bool {
	// 根据时间排序
	return this[i].CreationTimestamp.Time.After(this[j].CreationTimestamp.Time)
}
func (this GW) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}
