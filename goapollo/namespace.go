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

type ApolloConnConfig struct {
	AppId         string `json:"appId"`
	Cluster       string `json:"cluster"`
	NamespaceName string `json:"namespaceName"`
	ReleaseKey    string `json:"releaseKey"`
}

type ApolloConfig struct {
	ApolloConnConfig
	Configurations map[string]interface{} `json:"configurations"`
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

	// Add notificationId
	cluster.Notifications.Store(namespace, int64(-1))

	return nil
}

func (s *Server) UpdateNamespaceWithValue(appId string, clusterName string, namespace string, value string) error {
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
	if !exists {
		return errors.New("namespace not exists")
	}

	// Add config
	cluster.Namespaces.Store(namespace, &Namespace{Value: value})

	// increment notificationId
	notificationId, ok := cluster.Notifications.Load(namespace)
	if ok {
		cluster.Notifications.Store(namespace, notificationId.(int64)+1)
	} else {
		cluster.Notifications.Store(namespace, int64(-1))
	}

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
	if !ok {
		return false
	}
	cluster.Namespaces.Delete(namespace)

	// Delete NotificationId
	_, ok = cluster.Notifications.Load(namespace)
	if !ok {
		return false
	}
	cluster.Notifications.Delete(namespace)

	return true
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

func (s *Server) getNotificationId(appId, clusterName, namespace string) int64 {
	// Get app
	appInterface, ok := s.appMap.Load(appId)
	if !ok {
		return -1
	}
	app := appInterface.(*App)

	// Get cluster
	clusterInterface, ok := app.Clusters.Load(clusterName)
	if !ok {
		return -1
	}
	cluster := clusterInterface.(*Cluster)

	// Get NotificationId
	notificationId, ok := cluster.Notifications.Load(namespace)
	if !ok {
		return -1
	}

	return notificationId.(int64)
}

func (s *Server) getConfigfileJsonContent(c *gin.Context) {
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

	ac := ApolloConfig{
		ApolloConnConfig: ApolloConnConfig{
			AppId:         appId,
			Cluster:       clusterName,
			NamespaceName: namespace,
		},
		Configurations: nil,
	}

	if err := json.Unmarshal([]byte(configValue), &ac.Configurations); err != nil {
		c.AbortWithStatus(http.StatusFound)
		return
	}

	// Return config value
	c.JSON(http.StatusOK, ac)
}

func (s *Server) getNotificationsJsonContent(c *gin.Context) {
	appId := c.Query("appId")
	clusterName := c.Query("cluster")
	notifications := c.Query("notifications")

	notificationsList := make([]Notification, 0)
	err := json.Unmarshal([]byte(notifications), &notificationsList)
	if err != nil {
		c.AbortWithStatus(http.StatusNotModified)
		return
	}

	for i := 0; i < len(notificationsList); i++ {
		notificationId := s.getNotificationId(appId, clusterName, notificationsList[i].NamespaceName)
		if notificationsList[i].NotificationId < notificationId {
			notificationsList[i].NotificationId = notificationId
		}
	}

	c.JSON(http.StatusOK, notificationsList)
}
