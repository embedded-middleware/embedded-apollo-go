package goapollo

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestClusterMethods(t *testing.T) {
	s, err := NewServer(&ServerConfig{})
	require.NoError(t, err)

	// Add app
	err = s.AddApp("myApp")
	assert.NoError(t, err)

	// Add cluster
	err = s.AddCluster("myApp", "myCluster")
	assert.NoError(t, err)

	// Add duplicate cluster
	err = s.AddCluster("myApp", "myCluster")
	assert.Error(t, err)

	// Get cluster
	cluster, ok := s.GetCluster("myApp", "myCluster")
	assert.True(t, ok)
	assert.Equal(t, "myCluster", cluster.Name)

	// Get non-existent cluster
	cluster, ok = s.GetCluster("myApp", "nonExistentCluster")
	assert.False(t, ok)
	assert.Nil(t, cluster)

	// Get clusters
	clusters, ok := s.GetClusters("myApp")
	assert.True(t, ok)
	assert.Len(t, clusters, 1)
	assert.Equal(t, "myCluster", clusters[0].Name)

	// Delete cluster
	deleted := s.DeleteCluster("myApp", "myCluster")
	assert.True(t, deleted)

	// Delete non-existent cluster
	deleted = s.DeleteCluster("myApp", "nonExistentCluster")
	assert.False(t, deleted)

	// Get clusters after deletion
	clusters, ok = s.GetClusters("myApp")
	assert.True(t, ok)
	assert.Len(t, clusters, 0)

	// Delete app
	deleted = s.DeleteApp("myApp")
	assert.True(t, deleted)
}
