package leet48

import (
	"leetcode/infra"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestRotateImage(t *testing.T) {

	Convey("TestRotateImage", t, func() {
		Convey("成功", func() {
			arr0 := []int{1, 2, 3}
			arr1 := []int{4, 5, 6}
			arr2 := []int{7, 8, 9}
			sudu := [][]int{arr0, arr1, arr2}
			RotateImage(sudu)

			arr0 = []int{7, 4, 1}
			arr1 = []int{8, 5, 2}
			arr2 = []int{9, 6, 3}
			want := [][]int{arr0, arr1, arr2}

			if !infra.IntArr2DComp(sudu, want) {
				t.Errorf("failed")
			}

		})

	})

}
