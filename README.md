# Go SSH Tunnel package & tool

## Package

### Usage

```go
package main

import (
    sshTunnel "github.com/udondan/go-ssh-tunnel"
)

func main() {
    t := sshTunnel.New(context.Background(), 8080, "example.com", 80)

    // do something with the tunnel

    t.Close()
}
```

## Tool

### Installation

```bash
go install github.com/udondan/go-ssh-tunnel/ssh-tunnel
```

### Usage 

```bash
ssh-tunnel --local 8080 --host example.com --remote 80
```

Press <kbd>Ctrl</kbd>+<kbd>C</kbd> to close the tunnel.