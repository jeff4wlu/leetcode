package leet142

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSingleNum(t *testing.T) {

	Convey("TestSingleNum", t, func() {
		Convey("用例1", func() {

			want, got := LinkedLstCycle([]int{3, 2, 0, -4}, 1)

			if got != want {
				t.Errorf("failed")
			}

		})

		Convey("用例2", func() {

			want, got := LinkedLstCycle([]int{1}, -1)

			if got != want {
				t.Errorf("failed")
			}

		})

		Convey("用例3", func() {

			want, got := LinkedLstCycle([]int{1, 2}, 0)

			if got != want {
				t.Errorf("failed")
			}

		})

		Convey("用例4", func() {

			want, got := LinkedLstCycle([]int{1}, 0)
			if got != want {
				t.Errorf("failed")
			}

		})

	})

}
