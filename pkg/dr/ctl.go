package dr

import (
	"github.com/gin-gonic/gin"
	"github.com/xuliangTang/athena/athena"
)

type DestinationRuleCtl struct {
	DestinationRuleService *DestinationRuleService `inject:"-"`
}

func NewDestinationRuleCtl() *DestinationRuleCtl {
	return &DestinationRuleCtl{}
}

func (this *DestinationRuleCtl) list(ctx *gin.Context) any {
	ns := ctx.DefaultQuery("ns", "default")
	return this.DestinationRuleService.ListByNs(ns)
}

func (this *DestinationRuleCtl) Build(athena *athena.Athena) {
	athena.Handle("GET", "/destinationRules", this.list)
}
