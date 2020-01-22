package leet149

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMaxPointsLine(t *testing.T) {

	Convey("TestMaxPointsLine", t, func() {
		Convey("用例1", func() {

			in:=[][]int{
				{1,1},{2,2},{3,3},
			}

			got := MaxPointsLine(in)
			want := 3

			if got != want {
				t.Errorf("failed")
			}

		})

		Convey("用例2", func() {

			in:=[][]int{
				{1,1},{3,2},{5,3},{4,1},{2,3},{1,4},
			}

			got := MaxPointsLine(in)
			want := 4

			if got != want {
				t.Errorf("failed")
			}

		})



	})

}
