package leet39

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCombinationSum(t *testing.T) {

	Convey("TestCombinationSum", t, func() {
		Convey("用例1", func() {
			nums := []int{2, 3, 6, 7}

			res := CombinationSum(nums, 7)
			got := [][]int{{2, 2, 3}, {7}}

			if res == nil || !comp(res, got) {
				t.Errorf("failed")
			}

		})

		Convey("用例2", func() {
			nums := []int{2, 3, 5}

			res := CombinationSum(nums, 8)
			got := [][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}}

			if res == nil || !comp(res, got) {
				t.Errorf("failed")
			}

		})

		Convey("边界", func() {
			nums := []int{2, 3, 5}

			res := CombinationSum(nums, 1)
			//var got [][]int

			if len(res) != 0 {
				t.Errorf("failed")
			}

		})

	})

}

func comp(a, b [][]int) bool {

	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		isSame := false
		for j := 0; j < len(b); j++ {
			if compOneArr(a[i], b[j]) {
				isSame = true
				break
			}
		}
		if !isSame {
			return false
		}
	}
	return true
}

func compOneArr(a, b []int) bool {

	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
