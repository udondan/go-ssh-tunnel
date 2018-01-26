# Go SSH Tunnel package & tool

## Package

### Usage

```go
package main

import (
    "context"
    "github.com/udondan/go-ssh-tunnel"
)

func main() {
    ctx := context.Background()
    t := sshTunnel.New(ctx, 8080, "example.com", 80)

    // do something with the tunnel

    t.Close()
}
```

## Tool

### Installation

```bash
go install github.com/udondan/go-ssh-tunnel/cmd/ssh-tunnel
```

### Usage 

```bash
ssh-tunnel --local 8080 --host example.com --remote 80
```

Press <kbd>Ctrl</kbd>+<kbd>C</kbd> to close the tunnel.