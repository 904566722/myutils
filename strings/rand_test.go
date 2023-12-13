package strings

import (
	"fmt"
	"testing"
)

func TestRandStr(t *testing.T) {
	for i := 0; i < 20; i++ {
		fmt.Println(RandStr(10))
	}
}
