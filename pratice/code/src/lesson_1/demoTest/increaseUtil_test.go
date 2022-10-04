package demoTest

import (
	"fmt"
	"testing"
)

func TestIncrease(t *testing.T) {
	t.Log("Start testing increase!")
	increase(1, 2)
	var result = increase(1, 2)
	fmt.Println(result)
}
