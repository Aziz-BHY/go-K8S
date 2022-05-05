package config

import (
	"flag"
	"log"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

var Myclientset *kubernetes.Clientset

func GetClient() {
	kubeconfig := flag.String("kubeconfig", "./kubeconfig.yaml", "")
	flag.Parse()
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)

	if err != nil {
		log.Fatal(err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}
	Myclientset = clientset
}
