package leet37

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSudokuSolver(t *testing.T) {

	Convey("TestSudokuSolver", t, func() {
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
			SudokuSolver(sudu, 0, 0)

			arr0 = []int{5, 6, 1, 8, 4, 7, 9, 2, 3}
			arr1 = []int{3, 7, 9, 5, 2, 1, 6, 8, 4}
			arr2 = []int{4, 2, 8, 9, 6, 3, 1, 7, 5}
			arr3 = []int{6, 1, 3, 7, 8, 9, 5, 4, 2}
			arr4 = []int{7, 9, 4, 6, 5, 2, 3, 1, 8}
			arr5 = []int{8, 5, 2, 1, 3, 4, 7, 9, 6}
			arr6 = []int{9, 3, 5, 4, 7, 8, 2, 6, 1}
			arr7 = []int{1, 4, 6, 2, 9, 5, 8, 3, 7}
			arr8 = []int{2, 8, 7, 3, 1, 6, 4, 5, 9}
			want := [][]int{arr0, arr1, arr2, arr3, arr4, arr5, arr6, arr7, arr8}

			if !comp(sudu, want) {
				t.Errorf("failed")
			}

		})

	})

}

func comp(got, want [][]int) bool {

	for i := 0; i < len(got); i++ {
		for j := 0; j < len(got[0]); j++ {
			if got[i][j] != want[i][j] {
				return false
			}
		}
	}
	return true
}
