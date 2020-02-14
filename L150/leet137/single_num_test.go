package leet137

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSingleNum(t *testing.T) {

	Convey("TestSingleNum", t, func() {
		Convey("用例1", func() {

			got := SingleNum([]int{2, 2, 3, 2})
			want := 3

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}

		})

		Convey("用例2", func() {

			got := SingleNum([]int{0, 1, 0, 1, 0, 1, 99})
			want := 99

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}

		})

	})

	Convey("TestSingleNum1", t, func() {
		Convey("用例1", func() {

			got := SingleNum1([]int{2, 2, 3, 2})
			want := 3

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}

		})

		Convey("用例2", func() {

			got := SingleNum1([]int{0, 1, 0, 1, 0, 1, 99})
			want := 99

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}

		})

	})

	Convey("TestSingleNum2", t, func() {
		Convey("用例1", func() {

			got := SingleNum2([]int{2, 2, 3, 2})
			want := 3

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}

		})

		Convey("用例2", func() {

			got := SingleNum2([]int{0, 1, 0, 1, 0, 1, 99})
			want := 99

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}

		})

	})

}
