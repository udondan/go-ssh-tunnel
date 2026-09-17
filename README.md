# Go SSH Tunnel package & tool

> **⚠️ Archived (2026)** – This project is no longer maintained.
>
> It is a thin wrapper around the system `ssh` binary (`ssh -f -N -L`) and has been superseded by better options:
>
> - **In Go:** use [`golang.org/x/crypto/ssh`](https://pkg.go.dev/golang.org/x/crypto/ssh) for native in-process port forwarding (see [this guide](https://eli.thegreenplace.net/2022/ssh-port-forwarding-with-go/)), or a library built on it such as [`elliotchance/sshtunnel`](https://github.com/elliotchance/sshtunnel) or [`glycerine/sshego`](https://github.com/glycerine/sshego).
> - **On the command line:** `ssh -N -L 8080:127.0.0.1:80 example.com`, or [`autossh`](https://www.harding.motd.ca/autossh/) for automatic reconnects.
>
> The repository stays online so existing imports and forks keep working. No further changes will be made.

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
    if err := t.Open(); err != nil {
      panic(err)
    }
    defer t.Close()

    // do something with the tunnel
}
```

## Tool

### Installation

```
go install github.com/udondan/go-ssh-tunnel/cmd/ssh-tunnel
```

### Usage

```
ssh-tunnel --local 8080 --host example.com --remote 80
```

Press `Ctrl`+`C` to close the tunnel.

## License

MIT – see [LICENSE](LICENSE).
