package leet136

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSingleNum(t *testing.T) {

	Convey("TestSingleNum", t, func() {
		Convey("用例1", func() {

			got := SingleNum([]int{2, 2, 1})
			want := 1

			if got != want {
				t.Errorf("failed")
			}

		})

		Convey("用例2", func() {

			got := SingleNum([]int{4, 1, 2, 1, 2})
			want := 4

			if got != want {
				t.Errorf("failed")
			}

		})

	})

}
