package leet240

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSearchMatrix(t *testing.T) {

	Convey("TestSearchMatrix", t, func() {
		Convey("找到", func() {
			matrix := [][]int{
				{1, 4, 7, 11, 15},
				{2, 5, 8, 12, 19},
				{3, 6, 9, 16, 22},
				{10, 13, 14, 17, 24},
				{18, 21, 23, 26, 30},
			}

			got := SearchMatrix(matrix, 5)
			want := true

			if got != want {
				t.Errorf("failed")
			}

		})

		Convey("找不到", func() {
			matrix := [][]int{
				{1, 4, 7, 11, 15},
				{2, 5, 8, 12, 19},
				{3, 6, 9, 16, 22},
				{10, 13, 14, 17, 24},
				{18, 21, 23, 26, 30},
			}

			got := SearchMatrix(matrix, 20)
			want := false

			if got != want {
				t.Errorf("failed")
			}

		})

	})

}
