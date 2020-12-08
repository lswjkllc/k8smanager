package services

import (
	"context"
	"fmt"

	appsv1 "k8s.io/api/apps/v1"
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

func (ks K8SService) CreatePod(namespace string, pod *v1.Pod) (*v1.Pod, error) {
	kpod, err := ks.clientset.CoreV1().Pods(namespace).Create(context.TODO(), pod, metav1.CreateOptions{})
	return kpod, err
}

func (ks K8SService) ListDeployment(namespace string) (*appsv1.DeploymentList, error) {
	deployments, err := ks.clientset.AppsV1().Deployments(namespace).List(context.TODO(), metav1.ListOptions{})
	return deployments, err
}

func (ks K8SService) GetDeployment(namespace, deployment string) (*appsv1.Deployment, error) {
	kdeployment, err := ks.clientset.AppsV1().Deployments(namespace).Get(context.TODO(), deployment, metav1.GetOptions{})
	return kdeployment, err
}

func (ks K8SService) CreateDeployment(namespace string, deployment *appsv1.Deployment) (*appsv1.Deployment, error) {
	kdeployment, err := ks.clientset.AppsV1().Deployments(namespace).Create(context.TODO(), deployment, metav1.CreateOptions{})
	return kdeployment, err
}

func (ks K8SService) ListNamespace() (*v1.NamespaceList, error) {
	namespaces, err := ks.clientset.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	return namespaces, err
}

func (ks K8SService) GetNamespace(name string) (*v1.Namespace, error) {
	namespace, err := ks.clientset.CoreV1().Namespaces().Get(context.TODO(), name, metav1.GetOptions{})
	return namespace, err
}

func (ks K8SService) CreateNamespace(name *v1.Namespace) (*v1.Namespace, error) {
	namespace, err := ks.clientset.CoreV1().Namespaces().Create(context.TODO(), name, metav1.CreateOptions{})
	return namespace, err
}

func (ks K8SService) ListService(namespace string) (*v1.ServiceList, error) {
	services, err := ks.clientset.CoreV1().Services(namespace).List(context.TODO(), metav1.ListOptions{})
	return services, err
}

func (ks K8SService) GetService(namespace, name string) (*v1.Service, error) {
	service, err := ks.clientset.CoreV1().Services(namespace).Get(context.TODO(), name, metav1.GetOptions{})
	return service, err
}

func (ks K8SService) CreateService(namespace string, name *v1.Service) (*v1.Service, error) {
	service, err := ks.clientset.CoreV1().Services(namespace).Create(context.TODO(), name, metav1.CreateOptions{})
	return service, err
}
