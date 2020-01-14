package leet105

import (
	"fmt"
	"leetcode/L120/leet94"
	"leetcode/infra"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestConstructBT(t *testing.T) {

	Convey("TestConstructBT", t, func() {
		Convey("用例1", func() {

			tmp := ConstructBT([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7})
			got := leet94.BTInorderTra(tmp)
			want := []int{9, 3, 15, 20, 7}
			fmt.Println()
			infra.PrintTree(tmp, t)
			fmt.Println()

			if !infra.IntArrSeqCmp(got, want) {
				t.Errorf("failed")
			}

		})

	})

}
