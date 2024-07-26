package util

import "testing"

func TestUint256(t *testing.T) {
	num := BitMaxNumDecimal(256)
	t.Log(num)
	t.Log(len(num.String()))
}
