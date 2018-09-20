package kubetasks

import (
	"log"

	"github.com/maorfr/kube-tasks/pkg/utils"
	"github.com/maorfr/skbn/pkg/skbn"
)

// SimpleBackup performs backup
func SimpleBackup(namespace, selector, container, path, dst string, parallel int, tag string) (string, error) {
	log.Println("Backup started!")
	dstPrefix, dstPath := utils.SplitInTwo(dst, "://")

	log.Println("Getting clients")
	k8sClient, dstClient, err := skbn.GetClients("k8s", dstPrefix, "", dstPath)
	if err != nil {
		return "", err
	}

	log.Println("Getting pods")
	pods, err := utils.GetPods(k8sClient, namespace, selector)
	if err != nil {
		return "", err
	}

	log.Println("Calculating paths. This may take a while...")
	fromToPathsAllPods, err := utils.GetFromAndToPathsFromK8s(k8sClient, pods, namespace, container, tag, dstPath)
	if err != nil {
		return "", err
	}

	log.Println("Starting files copy")
	if err := skbn.PerformCopy(k8sClient, dstClient, "k8s", dstPrefix, fromToPathsAllPods, parallel); err != nil {
		return "", err
	}

	log.Println("All done!")
	return tag, nil
}
