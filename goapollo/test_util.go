package goapollo

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

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
