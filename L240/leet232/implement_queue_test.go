package leet232

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestQueue(t *testing.T) {

	Convey("TestQueue", t, func() {
		Convey("1", func() {
			queue := newQueue()
			queue.push(1)
			queue.push(2)

			got := queue.peek()
			want := 1
			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}
			got = queue.pop()
			want = 1
			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}

			gotB := queue.empty()
			wantB := false
			if gotB != wantB {
				t.Errorf("failed")
			}

		})

	})

}
