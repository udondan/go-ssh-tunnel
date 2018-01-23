package sshTunnel

import (
	"context"
	"fmt"
	"os/exec"
	"syscall"
	"time"
)

// Tunnel represents an ssh tunnel definition
// Additional to the public propperties it holds a channel which can be used to close the tunnel
type Tunnel struct {
	Host   string
	Local  int
	Remote int
	ctx    context.Context
	cancel context.CancelFunc
}

// Open opens an ssh tunnel
func (t Tunnel) Open() error {

	// error message from ssh currently unavailable.
	// https://github.com/golang/go/issues/10338

	//var stderr bytes.Buffer
	tunnelString := fmt.Sprintf("%d:127.0.0.1:%d", t.Local, t.Remote)
	args := []string{"-f", "-N", "-L", tunnelString, t.Host}
	cmd := exec.CommandContext(t.ctx, "ssh", args...)
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
	//cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		//errMsg := strings.Trim(stderr.String(), "\r\n ")
		//return errors.New(errMsg)
		return err
	}

	go func() {
		select {
		case <-t.ctx.Done():
			syscall.Kill(-cmd.Process.Pid, syscall.SIGKILL)
			return
		}
	}()

	return nil
}

// Close closes the ssh tunnel
func (t Tunnel) Close() {
	t.cancel()

	// wait for ssh background process to be killed
	// TODO: how can we get a signal that the ssh process stopped?
	time.Sleep(1 * time.Second)
}

// New create a new ssh tunnel
func New(ctx context.Context, local int, host string, remote int) *Tunnel {
	myCtx, cancel := context.WithCancel(ctx)
	return &Tunnel{Host: host, Local: local, Remote: remote, ctx: myCtx, cancel: cancel}
}
