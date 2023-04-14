package goapollo

import (
	"fmt"
	"testing"

	"github.com/apolloconfig/agollo/v4/storage"
	"github.com/stretchr/testify/require"
)

type testChangeListener struct {
	listenKey string
	newValue  string
	notify    chan struct{}
}

func (tc *testChangeListener) OnChange(event *storage.ChangeEvent) {
	change, ok := event.Changes[tc.listenKey]
	if !ok {
		return
	}
	tc.newValue = change.NewValue.(string)
	tc.notify <- struct{}{}
}

func (tc *testChangeListener) OnNewestChange(event *storage.FullChangeEvent) {

}

func setupGoapollo(t *testing.T) (*Server, int) {
	cfg := &ServerConfig{
		Port: 0,
	}
	srv, err := NewServer(cfg)
	require.NoError(t, err)

	// Start server
	port, err := srv.Start()
	require.NoError(t, err)
	fmt.Println("Server started on port", port)

	require.Greater(t, port, 0)
	return srv, port
}
