package leet174

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestDungeonGame(t *testing.T) {

	Convey("TestDungeonGame", t, func() {
		Convey("用例1", func() {

			dungeon := [][]int{
				{1, -3, 3},
				{0, -2, 0},
				{-3, -3, -3},
			}

			got := DungeonGame(dungeon)
			want := 3

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}

		})

		Convey("用例2", func() {

			dungeon := [][]int{
				{-2, -3, 3},
				{-5, -10, 1},
				{10, 30, -5},
			}

			got := DungeonGame(dungeon)
			want := 7

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}

		})

	})

}
