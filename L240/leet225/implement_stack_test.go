package leet225

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestStack(t *testing.T) {

	Convey("TestStack", t, func() {
		Convey("1", func() {
			stack := newStack()
			stack.push(1)
			stack.push(2)

			got := stack.top()
			want := 2
			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}
			got = stack.pop()
			want = 2
			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}

			gotB := stack.empty()
			wantB := false
			if gotB != wantB {
				t.Errorf("failed")
			}

		})

	})

}
