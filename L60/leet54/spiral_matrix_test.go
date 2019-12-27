package leet54

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSpiralMatrix(t *testing.T) {

	Convey("TestSpiralMatrix", t, func() {
		Convey("1", func() {
			arr0 := []int{1, 2, 3}
			arr1 := []int{4, 5, 6}
			arr2 := []int{7, 8, 9}
			matrix := [][]int{arr0, arr1, arr2}
			got := SpiralMatrix(matrix)

			want := []int{1, 2, 3, 6, 9, 8, 7, 4, 5}

			for i := 0; i < len(got); i++ {
				if got[i] != want[i] {
					t.Errorf("failed")
				}
			}

		})

		Convey("2", func() {
			arr0 := []int{1, 2, 3, 4}
			arr1 := []int{5, 6, 7, 8}
			arr2 := []int{9, 10, 11, 12}
			matrix := [][]int{arr0, arr1, arr2}
			got := SpiralMatrix(matrix)

			want := []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7}

			for i := 0; i < len(got); i++ {
				if got[i] != want[i] {
					t.Errorf("failed")
				}
			}

		})

	})

}
