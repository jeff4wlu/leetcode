package leet80

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestRemoveDuplicate(t *testing.T) {

	Convey("TestRemoveDuplicate", t, func() {
		Convey("用例1", func() {

			got := RemoveDuplicate([]int{1, 1, 1, 2, 2, 3})
			want := 5

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}

		})

		Convey("用例2", func() {

			got := RemoveDuplicate([]int{0, 0, 1, 1, 1, 1, 2, 3, 3})
			want := 7

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}

		})

	})

}
