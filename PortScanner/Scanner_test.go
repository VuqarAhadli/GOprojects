package main

import (
	"net"
	"sync"
	"testing"
	"sync/atomic"
	"time"
)

func TestGetOutboundIP(t *testing.T) {
	ip := GetOutboundIP()
	if ip == nil || ip.IsUnspecified() {
		t.Errorf("Expected valid IP, got %v", ip)
	}
}

func TestScanhost(t *testing.T) {
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	defer ln.Close()

	port := ln.Addr().(*net.TCPAddr).Port

	var openPorts int32
	var wg sync.WaitGroup
	wg.Add(1)

	go scanhost("127.0.0.1", port, 1, port, &wg, &openPorts)


	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()

	select {
	case <-done:
	case <-time.After(2 * time.Second):
		t.Fatal("scanhost timed out")
	}

	if atomic.LoadInt32(&openPorts) != 1 {
		t.Errorf("Expected 1 open port, got %d", openPorts)
	}
}