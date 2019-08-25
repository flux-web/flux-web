package controllers

import (
	"os"
	"fmt"
	"errors"
	"context"
	"io/ioutil"
	"encoding/json"
	
	"github.com/ghodss/yaml"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"

	"github.com/ericchiang/k8s"
    corev1 "github.com/ericchiang/k8s/apis/core/v1"
)

type NamespaceController struct {
	beego.Controller
}

func (this *NamespaceController) ListNamespaces() {
	var l = logs.GetLogger()

	client, err := loadClient()
    if err != nil {
        l.Panic(err.Error())
	}

    var namespacesList corev1.NamespaceList
    if err := client.List(context.Background(), "", &namespacesList); err != nil {
        l.Panic(err.Error())
	}
	
	var namespaces []string
	for _, namespace := range namespacesList.Items {
		namespaces = append(namespaces, *namespace.Metadata.Name)
	}

	namespacesJson, err := json.Marshal(namespaces)
    if err != nil {
        l.Panic("Cannot encode to JSON ", err)
    }

	this.Ctx.Output.Body(namespacesJson)
}

func loadClient() (*k8s.Client, error) {
	kubeconfigPath, err := getKubeconfigPath()
    if err != nil {
        return k8s.NewInClusterClient()
    }

	kubeconfig, err := ioutil.ReadFile(kubeconfigPath)
	if err != nil {
        return nil, fmt.Errorf("read kubeconfig: %v", err)
    }
	
    // Unmarshal YAML into a Kubernetes config object.
    var config k8s.Config
    if err := yaml.Unmarshal(kubeconfig, &config); err != nil {
        return nil, fmt.Errorf("unmarshal kubeconfig: %v", err)
    }
    return k8s.NewClient(&config)
}

func getKubeconfigPath() (string, error) {
    value := os.Getenv("KUBE_CONFIG_PATH")
    if len(value) == 0 {
        return "", errors.New("kubeconfig path env not found")
    }
    return value, nil
}
