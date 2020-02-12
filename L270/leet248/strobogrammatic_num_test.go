package leet248

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestStrobogrammaticNum(t *testing.T) {

	Convey("TestStrobogrammaticNum", t, func() {
		Convey("1", func() {

			got := StrobogrammaticNum("50", "100")
			want := 3

			if got != want {
				t.Errorf("failed")
			}

		})

	})

}
