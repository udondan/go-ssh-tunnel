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
	if err := t.Open(); err != nil {
		panic(err)
	}
    defer t.Close()
    
    // do something with the tunnel
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

## License

    MIT License

    Copyright (c) 2018 Daniel Schroeder

    Permission is hereby granted, free of charge, to any person obtaining a copy
    of this software and associated documentation files (the "Software"), to deal
    in the Software without restriction, including without limitation the rights
    to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
    copies of the Software, and to permit persons to whom the Software is
    furnished to do so, subject to the following conditions:

    The above copyright notice and this permission notice shall be included in all
    copies or substantial portions of the Software.

    THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
    IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
    FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
    AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
    LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
    OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
    SOFTWARE.```
