package sshTunnel

import (
	"fmt"
	"os/exec"
	"syscall"
)

// Tunnel represents an ssh tunnel definition
// Additional to the public propperties it holds a channel which can be used to close the tunnel
type Tunnel struct {
	Host     string
	Local    int
	Remote   int
	stopchan chan struct{}
}

// Open opens an ssh tunnel
func (t Tunnel) Open() error {

	// error message from ssh currently unavailable.
	// https://github.com/golang/go/issues/10338

	//var stderr bytes.Buffer
	tunnelString := fmt.Sprintf("%d:127.0.0.1:%d", t.Local, t.Remote)
	args := []string{"-f", "-N", "-L", tunnelString, t.Host}
	cmd := exec.Command("ssh", args...)
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
	//cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		//errMsg := strings.Trim(stderr.String(), "\r\n ")
		//return errors.New(errMsg)
		return err
	}

	go func() {
		for {
			select {
			case <-t.stopchan:
				syscall.Kill(-cmd.Process.Pid, syscall.SIGKILL)
				return
			}
		}
	}()

	return nil
}

// Close closes the ssh tunnel
func (t Tunnel) Close() {
	close(t.stopchan)
}

// New create a new ssh tunnel
func New(local int, host string, remote int) *Tunnel {
	stopchan := make(chan struct{})
	return &Tunnel{Host: host, Local: local, Remote: remote, stopchan: stopchan}
}
