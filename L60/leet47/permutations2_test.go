package leet47

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestPermutations2(t *testing.T) {

	Convey("TestPermutations2", t, func() {
		Convey("用例1", func() {
			nums := []int{1, 1, 2}

			res := Permutations2(nums)
			got := [][]int{
				{1, 1, 2},
				{1, 2, 1},
				{2, 1, 1},
			}

			if res == nil || !comp(res, got) {
				t.Errorf("failed")
			}

		})

		Convey("用例2", func() {
			nums := []int{1, 1, 2, 2}

			res := Permutations2(nums)
			got := [][]int{
				{1, 1, 2, 2},
				{1, 2, 1, 2},
				{2, 1, 1, 2},
				{1, 2, 2, 1},
				{2, 2, 1, 1},
				{2, 1, 2, 1},
			}

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
