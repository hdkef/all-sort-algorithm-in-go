package countingsort

import (
	"fmt"
	"testing"
)

func TestCountingSortDesc(t *testing.T) {
	testTable := []struct {
		Input    []int
		Expected []int
	}{
		{
			Input:    []int{5, 7, 2, 3, 8, 1, 4, 6},
			Expected: []int{8, 7, 6, 5, 4, 3, 2, 1},
		},
	}

	for i, v := range testTable {
		t.Run(fmt.Sprintf("tes descending no %d", i), func(t *testing.T) {
			CountingSortDesc(v.Input, 8)
			for i, k := range v.Expected {
				if k != v.Input[i] {
					t.Errorf("Fail, expected %d got %d", k, v.Input[i])
				}
			}
		})
	}
}

func TestCountingSortAsc(t *testing.T) {
	testTable := []struct {
		Input    []int
		Expected []int
	}{
		{
			Input:    []int{5, 7, 2, 3, 8, 1, 4, 6},
			Expected: []int{1, 2, 3, 4, 5, 6, 7, 8},
		},
	}

	for i, v := range testTable {
		t.Run(fmt.Sprintf("tes descending no %d", i), func(t *testing.T) {
			CountingSortAsc(v.Input, 8)
			for i, k := range v.Expected {
				if k != v.Input[i] {
					t.Errorf("Fail, expected %d got %d", k, v.Input[i])
				}
			}
		})
	}
}
