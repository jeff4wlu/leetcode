package leet217

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestContainDuplicate(t *testing.T) {

	Convey("TestContainDuplicate", t, func() {
		Convey("无重复", func() {

			got := ContainDuplicate([]int{1, 2})
			want := false

			if got != want {
				t.Errorf("failed")
			}

		})

		Convey("有重复1", func() {

			got := ContainDuplicate([]int{1, 2, 1})
			want := true

			if got != want {
				t.Errorf("failed")
			}

		})

		Convey("有重复2", func() {

			got := ContainDuplicate([]int{1, 1})
			want := true

			if got != want {
				t.Errorf("failed")
			}

		})

	})

}
