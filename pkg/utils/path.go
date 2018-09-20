package utils

import (
	"path/filepath"

	"github.com/maorfr/skbn/pkg/skbn"
)

// GetFromAndToPathsFromK8s aggregates paths from all pods
func GetFromAndToPathsFromK8s(iClient interface{}, pods []string, namespace, container, path, dstBasePath string) ([]skbn.FromToPair, error) {
	k8sClient := iClient.(*skbn.K8sClient)
	var fromToPathsAllPods []skbn.FromToPair
	for _, pod := range pods {

		fromToPaths, err := GetFromAndToPathsK8sToDst(k8sClient, namespace, pod, container, path, dstBasePath)
		if err != nil {
			return nil, err
		}
		fromToPathsAllPods = append(fromToPathsAllPods, fromToPaths...)
	}

	return fromToPathsAllPods, nil
}

// GetFromAndToPathsK8sToDst performs a path mapping between Kubernetes and a destination
func GetFromAndToPathsK8sToDst(k8sClient interface{}, namespace, pod, container, path, dstBasePath string) ([]skbn.FromToPair, error) {
	var fromToPaths []skbn.FromToPair

	pathPrfx := filepath.Join(namespace, pod, container, path)

	relativePaths, err := skbn.GetListOfFilesFromK8s(k8sClient, pathPrfx, "f", "*")
	if err != nil {
		return nil, err
	}

	for _, relativePath := range relativePaths {

		fromPath := filepath.Join(pathPrfx, relativePath)
		toPath := filepath.Join(dstBasePath, relativePath)

		fromToPaths = append(fromToPaths, skbn.FromToPair{FromPath: fromPath, ToPath: toPath})
	}

	return fromToPaths, nil
}
