package leet119

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestPascalTriangle(t *testing.T) {

	Convey("TestPascalTriangle", t, func() {
		Convey("用例1", func() {

			got := PascalTriangle(5, 3)

			fmt.Println()
			for i := 0; i < len(got); i++ {
				fmt.Printf("%d ", got[i])

			}
			fmt.Println()

		})

	})

}
