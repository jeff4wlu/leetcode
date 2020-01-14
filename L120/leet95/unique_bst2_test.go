package leet95

import (
	"fmt"
	"leetcode/infra"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestUniqueBST(t *testing.T) {

	Convey("TestUniqueBST", t, func() {
		Convey("用例1", func() {

			got := UniqueBST(3)
			fmt.Println()
			for _, v := range got {
				infra.PrintTree(v, t)
				fmt.Println()
			}

		})

	})

}
