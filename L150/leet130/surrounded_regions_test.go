package leet130

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSurroundedRegions(t *testing.T) {

	Convey("TestSurroundedRegions", t, func() {
		Convey("用例1", func() {

			board := [][]byte{
				{'o', 'o', 'x', 'x'},
				{'x', 'o', 'o', 'x'},
				{'x', 'x', 'o', 'x'},
				{'x', 'o', 'x', 'x'},
			}

			SurroundedRegions(&board)

			fmt.Println()
			for i := 0; i < len(board); i++ {
				for j := 0; j < len(board[0]); j++ {
					fmt.Printf("%q ", board[i][j])
				}
				fmt.Println()
			}

		})

		Convey("用例2", func() {

			board := [][]byte{
				{'o', 'x', 'x', 'x'},
				{'x', 'o', 'o', 'x'},
				{'x', 'x', 'o', 'x'},
				{'x', 'o', 'x', 'x'},
			}

			SurroundedRegions(&board)

			fmt.Println()
			for i := 0; i < len(board); i++ {
				for j := 0; j < len(board[0]); j++ {
					fmt.Printf("%q ", board[i][j])
				}
				fmt.Println()
			}

		})

	})

}
