package vs

import (
	"github.com/gin-gonic/gin"
	"github.com/xuliangTang/athena/athena"
	"istio.io/client-go/pkg/apis/networking/v1alpha3"
	"net/http"
)

type VsCtl struct {
	VsService *VsService `inject:"-"`
}

func NewVsCtl() *VsCtl {
	return &VsCtl{}
}

func (this *VsCtl) VsList(c *gin.Context) any {
	ns := c.DefaultQuery("ns", "default")

	return this.VsService.ListByNs(ns)
}

func (this *VsCtl) CreateVS(ctx *gin.Context) any {
	vs := &v1alpha3.VirtualService{}
	athena.Error(ctx.BindJSON(&vs))
	athena.Error(this.VsService.Create(vs))

	ctx.Set(athena.CtxHttpStatusCode, http.StatusCreated)
	return vs
}

func (this *VsCtl) ShowVs(ctx *gin.Context) any {
	uri := &VirtualServiceUri{}
	athena.Error(ctx.BindUri(&uri))

	return this.VsService.Load(uri.Namespace, uri.Name)
}

func (this *VsCtl) UpdateVs(ctx *gin.Context) (v athena.Void) {
	vs := &v1alpha3.VirtualService{}
	athena.Error(ctx.BindJSON(&vs))
	athena.Error(this.VsService.Update(vs))

	ctx.Set(athena.CtxHttpStatusCode, http.StatusNoContent)
	return
}

func (this *VsCtl) DeleteVs(ctx *gin.Context) (v athena.Void) {
	uri := &VirtualServiceUri{}
	athena.Error(ctx.BindUri(&uri))
	athena.Error(this.VsService.Delete(uri.Namespace, uri.Name))

	ctx.Set(athena.CtxHttpStatusCode, http.StatusNoContent)
	return
}

func (this *VsCtl) Build(athena *athena.Athena) {
	// 虚拟服务列表
	athena.Handle("GET", "/virtualServices", this.VsList)
	// 创建虚拟服务
	athena.Handle("POST", "/virtualService", this.CreateVS)
	// 虚拟服务详情
	athena.Handle("GET", "/virtualService/:ns/:name", this.ShowVs)
	// 更新虚拟服务
	athena.Handle("PUT", "/virtualService", this.UpdateVs)
	// 删除虚拟服务
	athena.Handle("DELETE", "/virtualService/:ns/:name", this.DeleteVs)
}
