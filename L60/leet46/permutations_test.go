package leet46

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestPermutations(t *testing.T) {

	Convey("TestPermutations", t, func() {
		Convey("用例1", func() {
			nums := []int{1, 2, 3}

			res := Permutations(nums)
			got := [][]int{
				{1, 2, 3},
				{1, 3, 2},
				{2, 1, 3},
				{2, 3, 1},
				{3, 1, 2},
				{3, 2, 1},
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
