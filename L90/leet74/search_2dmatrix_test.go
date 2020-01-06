package leet74

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSearch2DMatrix(t *testing.T) {

	Convey("TestSearch2DMatrix", t, func() {
		Convey("用例1", func() {
			r1 := []int{1, 3, 5, 7}
			r2 := []int{10, 11, 16, 20}
			r3 := []int{23, 30, 34, 50}

			got := Search2DMatrix([][]int{r1, r2, r3}, 3)
			want := true

			if got != want {
				t.Errorf("failed")
			}

		})

		Convey("用例2", func() {
			r1 := []int{1, 3, 5, 7}
			r2 := []int{10, 11, 16, 20}
			r3 := []int{23, 30, 34, 50}

			got := Search2DMatrix([][]int{r1, r2, r3}, 13)
			want := false

			if got != want {
				t.Errorf("failed")
			}

		})

		Convey("用例3", func() {
			r1 := []int{1, 3, 5, 7}
			r2 := []int{10, 11, 16, 20}
			r3 := []int{23, 30, 34, 50}

			got := Search2DMatrix([][]int{r1, r2, r3}, 2)
			want := false

			if got != want {
				t.Errorf("failed")
			}

		})
		Convey("用例4", func() {
			r1 := []int{1, 3, 5, 7}
			r2 := []int{10, 11, 16, 20}
			r3 := []int{23, 30, 34, 50}

			got := Search2DMatrix([][]int{r1, r2, r3}, 34)
			want := false

			if got != want {
				t.Errorf("failed")
			}

		})

	})

}
