package leet243

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestShortestDis(t *testing.T) {

	Convey("TestShortestDis", t, func() {
		Convey("1", func() {

			got := ShortestDis([]string{"practice", "makes", "perfect", "coding", "makes"}, "coding", "practice")
			want := 3

			if got != want {
				t.Errorf("failed")
			}

		})

		Convey("2", func() {

			got := ShortestDis([]string{"practice", "makes", "perfect", "coding", "makes"}, "coding", "makes")
			want := 1

			if got != want {
				t.Errorf("failed")
			}

		})

	})

}
