package dr

import (
	"github.com/gin-gonic/gin"
	"github.com/xuliangTang/athena/athena"
	"istio.io/client-go/pkg/apis/networking/v1alpha3"
	"net/http"
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

func (this *DestinationRuleCtl) delete(ctx *gin.Context) (v athena.Void) {
	uri := &DestinationRuleUri{}
	athena.Error(ctx.BindUri(&uri))
	athena.Error(this.DestinationRuleService.Delete(uri.Namespace, uri.Name))

	ctx.Set(athena.CtxHttpStatusCode, http.StatusNoContent)
	return
}

func (this *DestinationRuleCtl) show(ctx *gin.Context) any {
	uri := &DestinationRuleUri{}
	athena.Error(ctx.BindUri(&uri))

	return this.DestinationRuleService.Show(uri.Namespace, uri.Name)
}

func (this *DestinationRuleCtl) create(ctx *gin.Context) any {
	dr := &v1alpha3.DestinationRule{}
	athena.Error(ctx.BindJSON(dr))
	athena.Error(this.DestinationRuleService.Create(dr))

	ctx.Set(athena.CtxHttpStatusCode, http.StatusCreated)
	return dr
}

func (this *DestinationRuleCtl) update(ctx *gin.Context) (v athena.Void) {
	dr := &v1alpha3.DestinationRule{}
	athena.Error(ctx.BindJSON(dr))
	athena.Error(this.DestinationRuleService.Update(dr))

	ctx.Set(athena.CtxHttpStatusCode, http.StatusNoContent)
	return
}

func (this *DestinationRuleCtl) Build(athena *athena.Athena) {
	athena.Handle("GET", "/destinationRules", this.list)
	athena.Handle("DELETE", "/destinationRule/:ns/:name", this.delete)
	athena.Handle("GET", "/destinationRule/:ns/:name", this.show)
	athena.Handle("POST", "/destinationRule", this.create)
	athena.Handle("PUT", "/destinationRule", this.update)
}
