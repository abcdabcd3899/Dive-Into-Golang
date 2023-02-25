package main

import (
	"fmt"
	"testing"
	"time"
)

func TestMain(t *testing.T) {
	// setup test
	ch := make(chan int, 1)
	go func() {
		time.Sleep(time.Millisecond * 1)
		ch <- 1
	}()
	// select statement
	var output string
	select {
	case ch := <-ch:
		output = fmt.Sprintf("%d", ch)
	case <-time.After(time.Millisecond * 100):
		output = "Time out."
	default:
		output = "Receive nothing."
	}
	// assert output
	if output != "1" {
		t.Errorf("Expected output to be 1, but got %s", output)
	}
}
