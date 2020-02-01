package leet172

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestFactorialTailing(t *testing.T) {

	Convey("TestFactorialTailing", t, func() {
		Convey("用例1", func() {

			got := FactorialTailing(3)
			want := 0

			if got != want {
				t.Errorf("failed.")
			}

		})

		Convey("用例2", func() {

			got := FactorialTailing(16)
			want := 3

			if got != want {
				t.Errorf("failed.")
			}

		})

	})

}
