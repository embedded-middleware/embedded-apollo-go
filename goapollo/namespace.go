package goapollo

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Namespace struct {
	Value string
}

func (s *Server) AddNamespaceWithValue(appId string, clusterName string, namespace string, value string) error {
	// Get app
	appInterface, ok := s.appMap.Load(appId)
	if !ok {
		return errors.New("app does not exist")
	}
	app := appInterface.(*App)

	// Get cluster
	clusterInterface, ok := app.Clusters.Load(clusterName)
	if !ok {
		return errors.New("cluster does not exist")
	}
	cluster := clusterInterface.(*Cluster)

	// Check if namespace already exists
	_, exists := cluster.Namespaces.Load(namespace)
	if exists {
		return errors.New("namespace already exists")
	}

	// Add config
	cluster.Namespaces.Store(namespace, &Namespace{Value: value})

	return nil
}

func (s *Server) DeleteNamespace(appId string, clusterName string, namespace string) bool {
	// Get app
	appInterface, ok := s.appMap.Load(appId)
	if !ok {
		return false
	}
	app := appInterface.(*App)

	// Get cluster
	clusterInterface, ok := app.Clusters.Load(clusterName)
	if !ok {
		return false
	}
	cluster := clusterInterface.(*Cluster)

	// Delete config
	_, ok = cluster.Namespaces.Load(namespace)
	if ok {
		cluster.Namespaces.Delete(namespace)
		return true
	}

	return false
}

func (s *Server) GetConfig(appId string, clusterName string, namespace string) (string, bool) {
	// Get app
	appInterface, ok := s.appMap.Load(appId)
	if !ok {
		return "", false
	}
	app := appInterface.(*App)

	// Get cluster
	clusterInterface, ok := app.Clusters.Load(clusterName)
	if !ok {
		return "", false
	}
	cluster := clusterInterface.(*Cluster)

	// Get namespace
	namespaceInterface, ok := cluster.Namespaces.Load(namespace)
	if !ok {
		return "", false
	}
	namespaceValue := namespaceInterface.(*Namespace)

	// Get config value
	configValue := namespaceValue.Value

	if configValue == "" {
		return "", false
	}

	return configValue, true
}

func (s *Server) getConfigJsonContent(c *gin.Context) {
	appId := c.Param("appId")
	clusterName := c.Param("cluster")
	namespace := c.Param("namespace")

	// Get config value from map
	configValue, ok := s.GetConfig(appId, clusterName, namespace)
	if !ok {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	// Return config value
	c.JSON(http.StatusOK, json.RawMessage(configValue))
}
