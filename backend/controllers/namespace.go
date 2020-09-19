package controllers

import (
	"encoding/json"
	"os"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

type NamespaceController struct {
	beego.Controller
}

func (this *NamespaceController) ListNamespaces() {
	var l = logs.GetLogger()

	clientset := loadClient()

	namespacesList, err := clientset.CoreV1().Namespaces().List(metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}

	var namespaces []string
	for _, namespace := range namespacesList.Items {
		namespaces = append(namespaces, namespace.ObjectMeta.Name)
	}

	namespacesJson, err := json.Marshal(namespaces)
	if err != nil {
		l.Panic("Cannot encode to JSON ", err)
	}

	this.Ctx.Output.Body(namespacesJson)
}

func loadClient() *kubernetes.Clientset {
	kubeconfig, err := getKubeconfig()
	clientset, err := kubernetes.NewForConfig(kubeconfig)
	if err != nil {
		panic(err.Error())
	}
	return clientset
}

func getKubeconfig() (*rest.Config, error) {
	configPathFromEnv := os.Getenv("KUBE_CONFIG_PATH")

	var kubeconfig *rest.Config

	var err error
	if len(configPathFromEnv) != 0 {
		kubeconfig, err = clientcmd.BuildConfigFromFlags("", configPathFromEnv)
	} else {
		kubeconfig, err = rest.InClusterConfig()
	}

	if err != nil {
		panic(err.Error())
	}

	return kubeconfig, nil
}
