package leet108

import (
	"fmt"
	"leetcode/infra"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSortedArrBst(t *testing.T) {

	Convey("TestSortedArrBst", t, func() {
		Convey("用例1", func() {

			tmp := SortedArrBst([]int{-10, -3, 0, 5, 9})
			fmt.Println()
			infra.PrintTree(tmp, t)
			fmt.Println()

		})

	})

}
