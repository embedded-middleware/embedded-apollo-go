package goapollo

import (
	"fmt"
	"github.com/apolloconfig/agollo/v4"
	"github.com/apolloconfig/agollo/v4/env/config"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestOfficialAgolloTestGetConfig(t *testing.T) {
	server, port := setupGoapollo(t)
	err := server.AddApp("appId")
	require.NoError(t, err)
	err = server.AddCluster("appId", "cluster")
	require.NoError(t, err)
	err = server.AddNamespaceWithValue("appId", "cluster", "namespace", "{\"key\":\"value\"}")
	require.NoError(t, err)
	defer server.Close()
	c := &config.AppConfig{
		AppID:          "appId",
		Cluster:        "cluster",
		NamespaceName:  "namespace",
		IP:             fmt.Sprintf("http://localhost:%d", port),
		IsBackupConfig: false,
	}
	client, err := agollo.StartWithConfig(func() (*config.AppConfig, error) {
		return c, nil
	})
	if err != nil {
		require.NoError(t, err)
	}
	content := client.GetConfig("namespace").GetContent()
	require.Equal(t, "key=value\n", content)
	value := client.GetConfig("namespace").GetStringValue("key", "")
	require.Equal(t, "value", value)
}
