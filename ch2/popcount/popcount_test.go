package popcount

import (
	"fmt"
	"testing"
)

func TestPopcount(t *testing.T) {
	for i, v := range PC {
		fmt.Printf("PC[%d] = %v\n", i, v)
	}
	if Popcount(0) == 1 {
		t.Log("Popcount(0)=1\n")
	}
}
