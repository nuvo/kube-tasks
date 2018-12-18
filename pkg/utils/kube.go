package utils

import (
	"log"
	"os"
	"path/filepath"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

// GetReadyPods gets a list of all ready ports according to defined namespace and selector
func GetReadyPods(iClient interface{}, namespace, selector string) ([]string, error) {

	clientSet := GetClientSet()
	pods, err := clientSet.CoreV1().Pods(namespace).List(metav1.ListOptions{
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

// GetClientSet returns a kubernetes ClientSet
func GetClientSet() *kubernetes.Clientset {
	var kubeconfig string
	if kubeConfigPath := os.Getenv("KUBECONFIG"); kubeConfigPath != "" {
		kubeconfig = kubeConfigPath
	} else {
		kubeconfig = filepath.Join(os.Getenv("HOME"), ".kube", "config")
	}

	config, err := buildConfigFromFlags("", kubeconfig)
	if err != nil {
		log.Fatal(err.Error())
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err.Error())
	}

	return clientset
}

func buildConfigFromFlags(context, kubeconfigPath string) (*rest.Config, error) {
	return clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		&clientcmd.ClientConfigLoadingRules{ExplicitPath: kubeconfigPath},
		&clientcmd.ConfigOverrides{
			CurrentContext: context,
		}).ClientConfig()
}
