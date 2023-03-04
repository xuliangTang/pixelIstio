package bootstrap

import (
	istio "istio.io/client-go/pkg/clientset/versioned"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"log"
)

type K8sConfig struct{}

func NewK8sConfig() *K8sConfig {
	return &K8sConfig{}
}

func (this *K8sConfig) IstioRestClient() *istio.Clientset {
	client, err := istio.NewForConfig(this.K8sRestConfig())
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func (*K8sConfig) K8sRestConfig() *rest.Config {
	config, err := clientcmd.BuildConfigFromFlags("", "./resources/config")

	config.Insecure = true

	if err != nil {
		log.Fatal(err)
	}
	return config
}

// InitClient 初始化client-go客户端
func (this *K8sConfig) InitClient() *kubernetes.Clientset {
	c, err := kubernetes.NewForConfig(this.K8sRestConfig())
	if err != nil {
		log.Fatal(err)
	}
	return c
}
