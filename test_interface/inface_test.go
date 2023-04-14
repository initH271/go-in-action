package test_interface

import (
	"testing"
)

func TestInvoke(t *testing.T) {
	var c Cat
	invoke(&c)
	(&c).PrintInfo()
	c.PrintInfo()
}
