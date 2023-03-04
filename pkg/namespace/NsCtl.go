package namespace

import (
	"github.com/gin-gonic/gin"
	"github.com/xuliangTang/athena/athena"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type NsCtl struct {
	Client *kubernetes.Clientset `inject:"-"`
}

func NewNsCtl() *NsCtl {
	return &NsCtl{}
}

func (this *NsCtl) ListAll(c *gin.Context) any {
	list, err := this.Client.CoreV1().Namespaces().List(c, v1.ListOptions{})
	athena.Error(err)

	ret := make([]*NsModel, len(list.Items))
	for index, item := range list.Items {
		istio := false
		if _, ok := item.Labels["istio-injection"]; ok {
			istio = true
		}
		ret[index] = &NsModel{Name: item.Name, Istio: istio}
	}
	return ret
}

func (*NsCtl) Name() string {
	return "VsCtl"
}

func (this *NsCtl) Build(athena *athena.Athena) {
	athena.Handle("GET", "/nslist", this.ListAll)
}
