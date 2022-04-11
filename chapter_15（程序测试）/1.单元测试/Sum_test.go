package __单元测试

import (
	"testing"
)

func TestGetSum(t *testing.T) {
	result := Sum(1, 2)
	if result != 3 {
		t.Errorf("result is wrong")
		return
	}
	t.Log("result is right")
}
