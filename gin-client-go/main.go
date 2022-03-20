package main

import (
	"context"
	"flag"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"k8s.io/klog"
	"path/filepath"
)

func main() {
	var kubeConfig *string
	ctx := context.Background()
	if home := homedir.HomeDir(); home != "" {
		kubeConfig = flag.String("kubeConfig", filepath.Join(home, ".kube", "config"), "absolute path to the kubeConfig file")
	} else {
		kubeConfig = flag.String("kubeConfig", "", "absolute path to the kubeConfig file")

	}
	flag.Parse()
	config, err := clientcmd.BuildConfigFromFlags("", *kubeConfig)
	if err != nil {
		klog.Fatal(err)
		return
	}
	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		klog.Fatal(err)
		return
	}
	namespaceList, err := clientSet.CoreV1().Namespaces().List(ctx, metav1.ListOptions{})
	if err != nil {
		klog.Fatal(err)
		return
	}
	namespaces := namespaceList.Items
	for _, namespace := range namespaces {
		fmt.Println("name===>", namespace.Name, "status===>", string(namespace.Status.Phase))
	}
}
