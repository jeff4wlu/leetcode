package leet40

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCombinationSumUniq(t *testing.T) {

	Convey("TestCombinationSumUniq", t, func() {
		Convey("用例1", func() {
			nums := []int{10, 1, 2, 7, 6, 1, 5}

			res := CombinationSumUniq(nums, 8)
			got := [][]int{{1, 7}, {1, 2, 5}, {2, 6}, {1, 1, 6}}

			if res == nil || !comp(res, got) {
				t.Errorf("failed")
			}

		})

		Convey("用例2", func() {
			nums := []int{2, 5, 2, 1, 2}

			res := CombinationSumUniq(nums, 5)
			got := [][]int{{5}, {1, 2, 2}}

			if res == nil || !comp(res, got) {
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
