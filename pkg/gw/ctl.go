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

	return this.GwService.ListByNs(ns)
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

func (this *GatewayCtl) Build(athena *athena.Athena) {
	athena.Handle("GET", "/gateways", this.GwList)
	athena.Handle("POST", "/gateway", this.CreateGateway)
	athena.Handle("DELETE", "/gateway/:ns/:name", this.DeleteGW)
}
