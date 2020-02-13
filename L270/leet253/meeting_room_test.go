package leet253

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMeetingRoom(t *testing.T) {

	Convey("TestMeetingRoom", t, func() {
		Convey("1", func() {

			ranges := [][]int{
				{0, 30}, {5, 10}, {15, 20},
			}

			got := MeetingRoom(ranges)
			want := 2

			if got != want {
				t.Errorf("failed")
			}

		})

		Convey("2", func() {

			ranges := [][]int{
				{15, 20}, {5, 10}, {0, 3},
			}

			got := MeetingRoom(ranges)
			want := 1

			if got != want {
				t.Errorf("failed")
			}

		})

		Convey("3", func() {

			ranges := [][]int{
				{10, 20}, {0, 3}, {5, 15},
			}

			got := MeetingRoom(ranges)
			want := 2

			if got != want {
				t.Errorf("failed")
			}

		})

		Convey("4", func() {

			ranges := [][]int{
				{1, 10}, {3, 15}, {5, 20},
			}

			got := MeetingRoom(ranges)
			want := 3

			if got != want {
				t.Errorf("failed")
			}

		})

	})

}
