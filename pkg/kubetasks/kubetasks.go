package kubetasks

import (
	"fmt"
	"log"
	"path/filepath"
	"time"

	"github.com/maorfr/kube-tasks/pkg/utils"
	"github.com/maorfr/skbn/pkg/skbn"
)

// SimpleBackup performs backup
func SimpleBackup(namespace, selector, container, path, dst string, parallel int, tag string) (string, error) {
	log.Println("Backup started!")
	dstPrefix, dstPath := utils.SplitInTwo(dst, "://")
	dstPath = filepath.Join(dstPath, tag)

	log.Println("Getting clients")
	k8sClient, dstClient, err := skbn.GetClients("k8s", dstPrefix, "", dstPath)
	if err != nil {
		return "", err
	}

	log.Println("Getting pods")
	pods, err := utils.GetReadyPods(k8sClient, namespace, selector)
	if err != nil {
		return "", err
	}

	if len(pods) == 0 {
		return "", fmt.Errorf("No pods were found in namespace %s by selector %s", namespace, selector)
	}

	log.Println("Calculating paths. This may take a while...")
	fromToPathsAllPods, err := utils.GetFromAndToPathsFromK8s(k8sClient, pods, namespace, container, path, dstPath)
	if err != nil {
		return "", err
	}

	log.Println("Starting files copy to tag: " + tag)
	if err := skbn.PerformCopy(k8sClient, dstClient, "k8s", dstPrefix, fromToPathsAllPods, parallel); err != nil {
		return "", err
	}

	log.Println("All done!")
	return tag, nil
}

// WaitForPods waits for a given number of pods
func WaitForPods(namespace, selector string, desiredReplicas int) error {
	log.Println("Getting clients")
	k8sClient, err := skbn.GetClientToK8s()
	if err != nil {
		return err
	}

	readyPods := -1
	log.Printf("Waiting for %d ready pods", desiredReplicas)
	for readyPods != desiredReplicas {
		pods, err := utils.GetReadyPods(k8sClient, namespace, selector)
		if err != nil {
			return err
		}
		readyPods = len(pods)
		log.Printf("Currently %d ready pods", readyPods)
		if readyPods == desiredReplicas {
			break
		}
		time.Sleep(10 * time.Second)
	}
	return nil
}
