package services

import (
	"context"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

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
