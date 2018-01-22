package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	sshTunnel "github.com/udondan/go-ssh-tunnel"
)

func main() {

	local := flag.Int("local", 0, "The local port, e.g. 8080")
	host := flag.String("host", "", "The target host, e.g. example.com")
	remote := flag.Int("remote", 0, "The target port, e.g. 80")
	flag.Parse()

	if *local == 0 {
		fmt.Fprintln(os.Stderr, "Local port is required")
		flag.Usage()
		os.Exit(1)
	}
	if *host == "" {
		fmt.Fprintln(os.Stderr, "Target host is required")
		flag.Usage()
		os.Exit(1)
	}
	if *remote == 0 {
		fmt.Fprintln(os.Stderr, "Remote port is required")
		flag.Usage()
		os.Exit(1)
	}

	t := sshTunnel.New(*local, *host, *remote)

	err := t.Open()
	if err != nil {
		exitErr(err)
	}

	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		_ = <-sigs
		done <- true
	}()

	fmt.Println("Press Ctrl-C to close tunnel")
	<-done
	t.Close()
	time.Sleep(1 * time.Second)
}

func exitErr(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}
