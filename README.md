# embedded-apollo-go

Embedded Apollo for Go (embedded-apollo-go) is an easy to embedded apollo server, typically used in unit test.

## Installation

```bash
go get github.com/embedded-middleware/embedded-apollo-go
```

## Usage

Embedded Apollo for Go is typically used for unit testing, where you need to quickly spin up a mock server to handle
incoming requests. Here's an example of how to use it:

```go
package main

import (
	"fmt"

	"github.com/embedded-middleware/embedded-apollo-go/goapollo"
)

func main() {
	// Create server configuration
	cfg := &goapollo.ServerConfig{
		Port: 0,
	}

	// Create server instance
	server, err := goapollo.NewServer(cfg)
	if err != nil {
		// Handle error
	}

	// Start server
	port, err := server.Start()
	if err != nil {
		// Handle error
	}
	defer server.Close()

	// Use the server to handle incoming requests...

	// Print the allocated port
	fmt.Printf("Server started on port %d", port)
}
```

In this example, we create a new instance of the server using `goapollo.NewServer()`, start it using `server.Start()`, use it to handle incoming requests, and then shut it down using `server.Close()`.

Note that we set the `Port` field of the server configuration to 0, which tells the server to allocate a random port. We
also use the `defer` statement to ensure that the server is shut down when the main function exits.
