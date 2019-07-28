package cycle

import "testing"

func TestValidIntCycle1(t *testing.T) {
	validContent := []int64{
		0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2,
	}

	maxLoop := len(validContent)
	c := InitIntCycle(0, -1, 9)
	for i := 0; i < maxLoop; i++ {
		val := c.Cycle()
		if val != validContent[i] {
			t.Errorf("On loop %d expected %d, got %d", i, val, validContent[i])
		}
	}
}

func TestValidIntCycle2(t *testing.T) {
	validContent := []int64{
		2, 3, 4, 5, 6, 7, 8, 9, 2, 3, 4,
	}

	maxLoop := len(validContent)
	c := InitIntCycle(2, 1, 9)
	for i := 0; i < maxLoop; i++ {
		val := c.Cycle()
		if val != validContent[i] {
			t.Errorf("On loop %d expected %d, got %d", i, val, validContent[i])
		}
	}
}
