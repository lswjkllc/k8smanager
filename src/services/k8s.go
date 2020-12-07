package services

import (
	"context"
	"flag"
	"fmt"
	"time"

	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

(
	kubeconfig := flag.String("kubeconfig", "./config/config.yaml", "kubeconfig file path")
	flag.Parse()

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	// create the clientset: *kubernetes.Clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
)

func ListNode() {
	// get nodes
	nodes, err := clientset.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		fmt.Printf("%v", err)
		return
	}
	fmt.Printf("There are %d nodes in the cluster\n", len(nodes.Items))
	for i, node := range nodes.Items {
		fmt.Println(i, node.Name)
	}
}

func ListPod() {
	pods, err := clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))
	for i, pod := range pods.Items {
		fmt.Println(i, pod.Name)
	}
}

func GetPod(namespace, pod string) {
	_, err := clientset.CoreV1().Pods(namespace).Get(context.TODO(), pod, metav1.GetOptions{})
	if errors.IsNotFound(err) {
		fmt.Printf("Pod %s in namespace %s not found\n", pod, namespace)
	} else if statusError, isStatus := err.(*errors.StatusError); isStatus {
		fmt.Printf("Error getting pod %s in namespace %s: %v\n",
			pod, namespace, statusError.ErrStatus.Message)
	} else if err != nil {
		panic(err.Error())
	} else {
		fmt.Printf("Found pod {%s} in namespace {%s}\n", pod, namespace)
	}
}

func ConnectK8S(configPath string) {
	var kubeconfig *string
	kubeconfig = flag.String("kubeconfig", configPath, "kubeconfig file path")
	flag.Parse()

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	// nodes
	ListNode(clientset)

	// pods
	ListPod(clientset)

	// cycle get the pod
	for {
		namespace := "default"
		pod := "nginx"
		GetPod(clientset, namespace, pod)

		time.Sleep(3 * time.Second)
	}
}
