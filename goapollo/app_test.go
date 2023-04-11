package goapollo

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAppMethods(t *testing.T) {
	s, err := NewServer(&ServerConfig{})
	require.NoError(t, err)

	// Add app
	err = s.AddApp("myApp")
	assert.NoError(t, err)

	// Add duplicate app
	err = s.AddApp("myApp")
	assert.Error(t, err)

	// Get app
	app, ok := s.GetApp("myApp")
	assert.True(t, ok)
	assert.Equal(t, "myApp", app.AppId)

	// Get non-existent app
	app, ok = s.GetApp("nonExistentApp")
	assert.False(t, ok)
	assert.Nil(t, app)

	// Get apps
	apps := s.GetApps()
	assert.Len(t, apps, 1)
	assert.Equal(t, "myApp", apps[0].AppId)

	// Delete app
	deleted := s.DeleteApp("myApp")
	assert.True(t, deleted)

	// Delete non-existent app
	deleted = s.DeleteApp("nonExistentApp")
	assert.False(t, deleted)

	// Get apps after deletion
	apps = s.GetApps()
	assert.Len(t, apps, 0)
}
