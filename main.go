package main

import (
	"flag"
	"fmt"
	// "k8s.io/client-go/tools/clientcmd"
)

// out of cluster mode
var (
	kubeConfig = flag.String(
		"kubeconfig",
		"./config.yml",
		"absolute path to the kubelet config file")
)

func main() {
	fmt.Println("Beginning...")

	// config, err := clientcmd.BuildConfigFromFlags("", *kubeConfig)
	// if err != nil {
	// 	fmt.Println("Get k8s config failed %#v", err)
	// }
}
