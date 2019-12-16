package leet7

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestReverseInt(t *testing.T) {

	Convey("TestReverseInt", t, func() {
		Convey("常规测试", func() {
			in := -123
			got := ReverseInt(in)
			want := -321

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}

		})

		Convey("溢出测试", func() {
			in := 1000000009
			got := ReverseInt(in)
			var want int

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}

		})

		Convey("溢出边界测试--溢出", func() {
			in := 8463847412 //2147483647
			got := ReverseInt(in)
			want := 0

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}

		})

		Convey("溢出边界测试--无溢出", func() {
			in := 7463847412
			got := ReverseInt(in)
			want := 2147483647

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}

		})

	})

}
