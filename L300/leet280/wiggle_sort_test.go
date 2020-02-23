package leet280

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestWiggleSort(t *testing.T) {

	Convey("TestWiggleSort", t, func() {

		Convey("1", func() {

			got := WiggleSort([]int{3, 5, 2, 1, 6, 4})

			if !CmpWiggleSort(got) {
				t.Errorf("failed")
			}

		})

	})

}

func CmpWiggleSort(nums []int) bool {
	n := len(nums)
	if n < 2 {
		return false
	}

	for i := 1; i < n; i += 2 {
		if i-1 >= 0 {
			if nums[i] < nums[i-1] {
				return false
			}
		} else if i+1 < n {
			if nums[i] < nums[i+1] {
				return false
			}
		}
	}

	return true
}
