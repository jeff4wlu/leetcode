package leet249

import (
	"leetcode/infra"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGroupShifted(t *testing.T) {

	Convey("TestGroupShifted", t, func() {
		Convey("1", func() {

			got := GroupShifted([]string{"abc", "bcd", "acef", "xyz", "az", "ba", "a", "z"})
			want := [][]string{
				{"abc", "bcd", "xyz"},
				{"az", "ba"},
				{"acef"},
				{"a", "z"},
			}

			if !infra.StringArrCollectionComp(got, want) {
				t.Errorf("failed")
			}

		})

	})

}
