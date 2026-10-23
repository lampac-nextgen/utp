package utp

import (
	"testing"
	"time"
)

func TestReadFromReturnsWhenSocketClosed(t *testing.T) {
	s, err := NewSocket("udp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	done := make(chan error, 1)
	go func() {
		buf := make([]byte, 1500)
		_, _, err := s.ReadFrom(buf)
		done <- err
	}()
	if err := s.Close(); err != nil {
		t.Fatal(err)
	}
	select {
	case err := <-done:
		if err == nil {
			t.Fatal("ReadFrom returned nil error after Close")
		}
	case <-time.After(2 * time.Second):
		t.Fatal("ReadFrom blocked after Close")
	}
}
