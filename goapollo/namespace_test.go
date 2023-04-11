package goapollo

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNamespaceMethods(t *testing.T) {
	s, err := NewServer(&ServerConfig{})
	require.NoError(t, err)

	// Add app
	err = s.AddApp("myApp")
	assert.NoError(t, err)

	// Add cluster
	err = s.AddCluster("myApp", "myCluster")
	assert.NoError(t, err)

	// Add config
	err = s.AddNamespaceWithValue("myApp", "myCluster", "myNamespace", "myConfig")
	assert.NoError(t, err)

	// Get config
	config, ok := s.GetConfig("myApp", "myCluster", "myNamespace")
	assert.True(t, ok)
	assert.Equal(t, "myConfig", config)

	// Get non-existent config
	config, ok = s.GetConfig("myApp", "myCluster", "nonExistentNamespace")
	assert.False(t, ok)
	assert.Equal(t, "", config)

	// Delete config
	deleted := s.DeleteNamespace("myApp", "myCluster", "myNamespace")
	assert.True(t, deleted)

	// Delete non-existent config
	deleted = s.DeleteNamespace("myApp", "myCluster", "nonExistentNamespace")
	assert.False(t, deleted)

	// Get config after deletion
	config, ok = s.GetConfig("myApp", "myCluster", "myNamespace")
	assert.False(t, ok)
	assert.Equal(t, "", config)
}
