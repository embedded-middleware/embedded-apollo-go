package goapollo

import (
	"errors"
	"sync"
)

type Cluster struct {
	Name       string
	Namespaces sync.Map
}

func (s *Server) AddCluster(appId string, clusterName string) error {
	// Get app
	appInterface, ok := s.appMap.Load(appId)
	if !ok {
		return errors.New("app does not exist")
	}
	app := appInterface.(*App)

	// Check if cluster already exists
	if _, ok := app.Clusters.Load(clusterName); ok {
		return errors.New("cluster already exists")
	}

	// Add cluster
	app.Clusters.Store(clusterName, &Cluster{Name: clusterName})

	return nil
}

func (s *Server) DeleteCluster(appId string, clusterName string) bool {
	// Get app
	appInterface, ok := s.appMap.Load(appId)
	if !ok {
		return false
	}
	app := appInterface.(*App)

	// Delete cluster
	_, ok = app.Clusters.LoadAndDelete(clusterName)

	return ok
}

func (s *Server) GetClusters(appId string) ([]*Cluster, bool) {
	// Get app
	appInterface, ok := s.appMap.Load(appId)
	if !ok {
		return nil, false
	}
	app := appInterface.(*App)

	// Get cluster list
	clusters := make([]*Cluster, 0)
	app.Clusters.Range(func(key, value interface{}) bool {
		clusters = append(clusters, value.(*Cluster))
		return true
	})

	return clusters, true
}

func (s *Server) GetCluster(appId string, clusterName string) (*Cluster, bool) {
	// Get app
	appInterface, ok := s.appMap.Load(appId)
	if !ok {
		return nil, false
	}
	app := appInterface.(*App)

	// Get cluster
	clusterInterface, ok := app.Clusters.Load(clusterName)
	if !ok {
		return nil, false
	}
	cluster := clusterInterface.(*Cluster)

	return cluster, true
}
