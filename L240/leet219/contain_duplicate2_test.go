package leet219

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestContainDuplicate2(t *testing.T) {

	Convey("TestContainDuplicate2", t, func() {
		Convey("1", func() {

			got := ContainDuplicate2([]int{1, 2, 3, 1}, 3)
			want := true

			if got != want {
				t.Errorf("failed")
			}

		})

		Convey("2", func() {

			got := ContainDuplicate2([]int{1, 0, 1, 1}, 1)
			want := true

			if got != want {
				t.Errorf("failed")
			}

		})

		Convey("3", func() {

			got := ContainDuplicate2([]int{1, 2, 3, 1, 2, 3}, 2)
			want := false

			if got != want {
				t.Errorf("failed")
			}

		})

	})

}
