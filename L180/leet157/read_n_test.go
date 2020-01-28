package leet157

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMaxPro(t *testing.T) {

	Convey("TestReadN", t, func() {
		Convey("1", func() {
			f := []byte{'a', 'b', 'c'}
			got := ReadN(f, 6)
			want := 3

			if got != want {
				t.Errorf("got %d, want  %d", got, want)
			}

		})

		Convey("不超过文件长度，但在一个循环外", func() {
			f := []byte{'a', 'b', 'c', 'a', 'b', 'c', 'a', 'b', 'c', 'a', 'b', 'c'}
			got := ReadN(f, 6)
			want := 6

			if got != want {
				t.Errorf("got %d, want  %d", got, want)
			}

		})

		Convey("超过文件长度", func() {
			f := []byte{'a', 'b', 'c', 'a', 'b', 'c', 'a', 'b', 'c', 'a', 'b', 'c'}
			got := ReadN(f, 20)
			want := 12

			if got != want {
				t.Errorf("got %d, want  %d", got, want)
			}

		})

		Convey("空文件", func() {
			f := []byte{}
			got := ReadN(f, 20)
			want := 0

			if got != want {
				t.Errorf("got %d, want  %d", got, want)
			}

		})

	})

}
