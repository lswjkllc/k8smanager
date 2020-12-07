package services

import (
	"context"
	"fmt"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

type K8SService struct {
	clientset *kubernetes.Clientset
}

func New() K8SService {
	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", "./config/config.yaml")
	if err != nil {
		panic(err.Error())

	}
	// create the clientset: *kubernetes.Clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	ks := K8SService{clientset}

	return ks
}

func (ks K8SService) ListNode() {
	// get nodes
	nodes, err := ks.clientset.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		fmt.Printf("%v", err)
		return
	}
	fmt.Printf("There are %d nodes in the cluster\n", len(nodes.Items))
	for i, node := range nodes.Items {
		fmt.Println(i, node.Name)
	}
}

func (ks K8SService) ListPod(namespace string) (*v1.PodList, error) {
	pods, err := ks.clientset.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{})
	return pods, err
}

func (ks K8SService) GetPod(namespace, pod string) (*v1.Pod, error) {
	kpod, err := ks.clientset.CoreV1().Pods(namespace).Get(context.TODO(), pod, metav1.GetOptions{})
	return kpod, err
}
