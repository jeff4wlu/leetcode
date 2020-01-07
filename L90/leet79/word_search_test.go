package leet79

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestWordSearch(t *testing.T) {

	Convey("TestWordSearch", t, func() {
		Convey("用例1", func() {

			board := [][]byte{
				[]byte{'A', 'B', 'C', 'E'},
				[]byte{'S', 'F', 'C', 'S'},
				[]byte{'A', 'D', 'E', 'E'},
			}

			got := WordSearch(board, "ABCCED")
			want := true

			if got != want {
				t.Errorf("got %v, want %v", got, want)
			}

		})

		Convey("用例2", func() {

			board := [][]byte{
				[]byte{'A', 'B', 'C', 'E'},
				[]byte{'S', 'F', 'C', 'S'},
				[]byte{'A', 'D', 'E', 'E'},
			}

			got := WordSearch(board, "ABCB")
			want := false

			if got != want {
				t.Errorf("got %v, want %v", got, want)
			}

		})

	})

}
