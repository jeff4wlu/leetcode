package leet60

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestPermutationSeq(t *testing.T) {

	Convey("TestPermutationSeq", t, func() {
		Convey("用例1", func() {
			got := PermutationSeq(4, 9)
			want := "2314"

			if got != want {
				t.Errorf("got %s, want %s", got, want)
			}

		})

	})

	Convey("TestPermutationSeq", t, func() {
		Convey("用例2", func() {
			got := PermutationSeq(4, 6)
			want := "1432"

			if got != want {
				t.Errorf("got %s, want %s", got, want)
			}

		})

	})

	Convey("TestPermutationSeq", t, func() {
		Convey("用例3", func() {
			got := PermutationSeq(3, 4)
			want := "231"

			if got != want {
				t.Errorf("got %s, want %s", got, want)
			}

		})

	})

}
