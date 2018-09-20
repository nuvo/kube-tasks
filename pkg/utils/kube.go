package utils

import (
	"github.com/maorfr/skbn/pkg/skbn"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func GetReadyPods(iClient interface{}, namespace, selector string) ([]string, error) {

	k8sClient := *iClient.(*skbn.K8sClient)
	pods, err := k8sClient.ClientSet.CoreV1().Pods(namespace).List(metav1.ListOptions{
		LabelSelector: selector,
	})
	if err != nil {
		return nil, err
	}

	var podList []string
	for _, pod := range pods.Items {
		ready := true
		for _, condition := range pod.Status.Conditions {
			if condition.Status != "True" {
				ready = false
				break
			}
		}
		if ready {
			podList = append(podList, pod.Name)
		}
	}

	return podList, nil
}
