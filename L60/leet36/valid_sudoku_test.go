package leet36

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestValidSudoku(t *testing.T) {

	Convey("TestValidSudoku", t, func() {
		Convey("成功", func() {
			arr0 := []int{5, 6, 0, 8, 4, 7, 0, 0, 0}
			arr1 := []int{3, 0, 9, 0, 0, 0, 6, 0, 0}
			arr2 := []int{0, 0, 8, 0, 0, 0, 0, 0, 0}
			arr3 := []int{0, 1, 0, 0, 8, 0, 0, 4, 0}
			arr4 := []int{7, 9, 0, 6, 0, 2, 0, 1, 8}
			arr5 := []int{0, 5, 0, 0, 3, 0, 0, 9, 0}
			arr6 := []int{0, 0, 0, 0, 0, 0, 2, 0, 0}
			arr7 := []int{0, 0, 6, 0, 0, 0, 8, 0, 7}
			arr8 := []int{0, 0, 0, 3, 1, 6, 0, 5, 9}
			sudu := [][]int{arr0, arr1, arr2, arr3, arr4, arr5, arr6, arr7, arr8}
			got := ValidSudoku(sudu)
			want := true

			if got != want {
				t.Errorf("got %v, want  %v", got, want)
			}

		})

		Convey("失败", func() {
			arr0 := []int{5, 3, 0, 0, 7, 0, 0, 0, 0}
			arr1 := []int{6, 0, 0, 1, 9, 5, 0, 0, 0}
			arr2 := []int{-1, 9, 8, 0, 0, 0, 0, 6, 0}
			arr3 := []int{8, 0, 0, 0, 6, 0, 0, 0, 3}
			arr4 := []int{4, 0, 0, 8, 0, 3, 0, 0, 1}
			arr5 := []int{7, 0, 0, 0, 2, 0, 0, 0, 6}
			arr6 := []int{-1, 6, 0, 0, 0, 0, 2, 8, 0}
			arr7 := []int{-1, 0, 0, 4, 1, 9, 0, 0, 5}
			arr8 := []int{-1, 0, 6, 0, 8, 0, 0, 7, 9}
			sudu := [][]int{arr0, arr1, arr2, arr3, arr4, arr5, arr6, arr7, arr8}
			got := ValidSudoku(sudu)
			want := false

			if got != want {
				t.Errorf("got %v, want  %v", got, want)
			}

		})

	})

}
