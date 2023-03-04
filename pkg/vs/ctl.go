package vs

import (
	"github.com/gin-gonic/gin"
	"github.com/xuliangTang/athena/athena"
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

func (this *VsCtl) Build(athena *athena.Athena) {
	athena.Handle("GET", "/virtualServices", this.VsList)
}
