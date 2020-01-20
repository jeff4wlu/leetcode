package leet141

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestLinkedLstCycle(t *testing.T) {

	Convey("TestLinkedLstCycle", t, func() {
		Convey("用例1", func() {

			got := LinkedLstCycle([]int{3, 2, 0, -4}, 1)
			want := true

			if got != want {
				t.Errorf("failed")
			}

		})

		Convey("用例2", func() {

			got := LinkedLstCycle([]int{1}, -1)
			want := false

			if got != want {
				t.Errorf("failed")
			}

		})

		Convey("用例3", func() {

			got := LinkedLstCycle([]int{1, 2}, 0)
			want := true

			if got != want {
				t.Errorf("failed")
			}

		})

		Convey("用例4", func() {

			got := LinkedLstCycle([]int{1}, 0)
			want := true

			if got != want {
				t.Errorf("failed")
			}

		})

	})

}
