package main_test

import (
	fib "abcdabcd3899/package_demo/test"
	"testing"
)

func TestPC(t *testing.T) {
	if v, err := fib.Fib(-10); err != nil {
		t.Log(v)
	}

}
