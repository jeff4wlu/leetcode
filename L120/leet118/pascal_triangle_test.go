package leet118

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestPascalTriangle(t *testing.T) {

	Convey("TestPascalTriangle", t, func() {
		Convey("用例1", func() {

			got := PascalTriangle(5)

			fmt.Println()
			for i := 0; i < len(got); i++ {
				for j := 0; j < len(got[i]); j++ {
					fmt.Printf("%d ", got[i][j])
				}
				fmt.Println()
			}

		})

	})

}
