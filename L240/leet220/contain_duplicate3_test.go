package leet220

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestContainDuplicate3(t *testing.T) {

	Convey("TestContainDuplicate3", t, func() {
		Convey("1", func() {

			got := ContainDuplicate3([]int{1, 2, 3, 1}, 0, 3)
			want := true

			if got != want {
				t.Errorf("failed")
			}

		})

		Convey("2", func() {

			got := ContainDuplicate3([]int{1, 0, 1, 1}, 2, 1)
			want := true

			if got != want {
				t.Errorf("failed")
			}

		})

		Convey("3", func() {

			got := ContainDuplicate3([]int{1, 5, 9, 1, 5, 9}, 3, 2)
			want := false

			if got != want {
				t.Errorf("failed")
			}

		})

	})

}
