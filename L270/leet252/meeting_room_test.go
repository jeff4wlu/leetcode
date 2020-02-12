package leet252

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
			want := false

			if got != want {
				t.Errorf("failed")
			}

		})

		Convey("2", func() {

			ranges := [][]int{
				{0, 3}, {5, 10}, {15, 20},
			}

			got := MeetingRoom(ranges)
			want := true

			if got != want {
				t.Errorf("failed")
			}

		})

		Convey("3", func() {

			ranges := [][]int{
				{0, 3}, {5, 15}, {10, 20},
			}

			got := MeetingRoom(ranges)
			want := false

			if got != want {
				t.Errorf("failed")
			}

		})

	})

}
