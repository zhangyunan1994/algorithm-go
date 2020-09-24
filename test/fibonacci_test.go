package test

import (
	"algorithm/other"
	"testing"
)

func TestFibonacci(t *testing.T)  {
	for i := 0; i < 10; i++ {
		println(other.Fib(i))
	}
	t.Log("success")
}
