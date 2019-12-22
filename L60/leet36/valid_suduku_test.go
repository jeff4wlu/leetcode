package leet36

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestValidSuduku(t *testing.T) {

	Convey("TestValidSuduku", t, func() {
		Convey("成功", func() {
			arr0 := []int{5, 3, -1, -1, 7, -1, -1, -1, -1}
			arr1 := []int{6, -1, -1, 1, 9, 5, -1, -1, -1}
			arr2 := []int{-1, 9, 8, -1, -1, -1, -1, 6, -1}
			arr3 := []int{8, -1, -1, -1, 6, -1, -1, -1, 3}
			arr4 := []int{4, -1, -1, 8, -1, 3, -1, -1, 1}
			arr5 := []int{7, -1, -1, -1, 2, -1, -1, -1, 6}
			arr6 := []int{-1, 6, -1, -1, -1, -1, 2, 8, -1}
			arr7 := []int{-1, -1, -1, 4, 1, 9, -1, -1, 5}
			arr8 := []int{-1, -1, -1, -1, 8, -1, -1, 7, 9}
			sudu := [][]int{arr0, arr1, arr2, arr3, arr4, arr5, arr6, arr7, arr8}
			got := ValidSuduku(sudu)
			want := true

			if got != want {
				t.Errorf("got %v, want  %v", got, want)
			}

		})

		Convey("失败", func() {
			arr0 := []int{5, 3, -1, -1, 7, -1, -1, -1, -1}
			arr1 := []int{6, -1, -1, 1, 9, 5, -1, -1, -1}
			arr2 := []int{-1, 9, 8, -1, -1, -1, -1, 6, -1}
			arr3 := []int{8, -1, -1, -1, 6, -1, -1, -1, 3}
			arr4 := []int{4, -1, -1, 8, -1, 3, -1, -1, 1}
			arr5 := []int{7, -1, -1, -1, 2, -1, -1, -1, 6}
			arr6 := []int{-1, 6, -1, -1, -1, -1, 2, 8, -1}
			arr7 := []int{-1, -1, -1, 4, 1, 9, -1, -1, 5}
			arr8 := []int{-1, -1, 6, -1, 8, -1, -1, 7, 9}
			sudu := [][]int{arr0, arr1, arr2, arr3, arr4, arr5, arr6, arr7, arr8}
			got := ValidSuduku(sudu)
			want := false

			if got != want {
				t.Errorf("got %v, want  %v", got, want)
			}

		})

	})

}
