package dr

import (
	"github.com/gin-gonic/gin"
	"istio.io/client-go/pkg/apis/networking/v1alpha3"
	"log"
	"pixelIstio/pkg/wscore"
)

type DestinationRuleHandler struct {
	DestinationRuleMap     *DestinationRuleMap     `inject:"-"`
	DestinationRuleService *DestinationRuleService `inject:"-"`
}

func (this *DestinationRuleHandler) OnAdd(obj interface{}) {
	this.DestinationRuleMap.Add(obj.(*v1alpha3.DestinationRule))
	ns := obj.(*v1alpha3.Gateway).Namespace
	wscore.ClientMap.SendAll(
		gin.H{
			"type": "dr",
			"result": gin.H{
				"ns":   ns,
				"data": this.DestinationRuleService.ListByNs(ns),
			},
		},
	)
}
func (this *DestinationRuleHandler) OnUpdate(oldObj, newObj interface{}) {
	err := this.DestinationRuleMap.Update(newObj.(*v1alpha3.DestinationRule))
	if err != nil {
		log.Println(err)
		return
	}
	ns := newObj.(*v1alpha3.Gateway).Namespace
	wscore.ClientMap.SendAll(
		gin.H{
			"type": "dr",
			"result": gin.H{
				"ns":   ns,
				"data": this.DestinationRuleService.ListByNs(ns),
			},
		},
	)
}

func (this *DestinationRuleHandler) OnDelete(obj interface{}) {
	this.DestinationRuleMap.Delete(obj.(*v1alpha3.DestinationRule))
	ns := obj.(*v1alpha3.Gateway).Namespace

	wscore.ClientMap.SendAll(
		gin.H{
			"type": "dr",
			"result": gin.H{
				"ns":   ns,
				"data": this.DestinationRuleService.ListByNs(ns),
			},
		},
	)
}
