package crd

import (
	apiextensions "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
)

const (
	clusterDefinitionFilename = "artifacts/mysql-crd.yaml"
)

// CreateClusterCRD registers a MySQLCluster custom resource in kubernetes api.
func CreateClusterCRD(clientset *apiextensions.Clientset) error {
	return createCRD(clientset, clusterDefinitionFilename)
}