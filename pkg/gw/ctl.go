package gw

import (
	"github.com/gin-gonic/gin"
	"github.com/xuliangTang/athena/athena"
	"istio.io/client-go/pkg/apis/networking/v1alpha3"
	"net/http"
)

type GatewayCtl struct {
	GwService *GateWayService `inject:"-"`
}

func NewGatewayCtl() *GatewayCtl {
	return &GatewayCtl{}
}

func (this *GatewayCtl) GwList(c *gin.Context) any {
	ns := c.DefaultQuery("ns", "default")
	if ns != "" {
		return this.GwService.ListByNs(ns)
	}

	return this.GwService.ListAll()
}

func (this *GatewayCtl) CreateGateway(c *gin.Context) any {
	gw := &v1alpha3.Gateway{}
	athena.Error(c.BindJSON(gw))
	athena.Error(this.GwService.Create(gw))

	c.Set(athena.CtxHttpStatusCode, http.StatusCreated)
	return gw
}

func (this *GatewayCtl) DeleteGW(c *gin.Context) (v athena.Void) {
	uri := &GatewayUri{}
	athena.Error(c.BindUri(&uri))
	athena.Error(this.GwService.Delete(uri.Namespace, uri.Name))

	c.Set(athena.CtxHttpStatusCode, http.StatusNoContent)
	return
}

func (this *GatewayCtl) ShowGW(ctx *gin.Context) any {
	uri := &GatewayUri{}
	athena.Error(ctx.BindUri(&uri))

	return this.GwService.Show(uri.Namespace, uri.Name)
}

func (this *GatewayCtl) Update(ctx *gin.Context) (v athena.Void) {
	gw := &v1alpha3.Gateway{}
	athena.Error(ctx.BindJSON(&gw))
	athena.Error(this.GwService.Update(gw))

	ctx.Set(athena.CtxHttpStatusCode, http.StatusNoContent)
	return
}

func (this *GatewayCtl) Build(athena *athena.Athena) {
	athena.Handle("GET", "/gateways", this.GwList)
	athena.Handle("POST", "/gateway", this.CreateGateway)
	athena.Handle("DELETE", "/gateway/:ns/:name", this.DeleteGW)
	athena.Handle("GET", "/gateway/:ns/:name", this.ShowGW)
	athena.Handle("PUT", "/gateway", this.Update)
}
