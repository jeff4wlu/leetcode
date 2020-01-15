package leet116

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestPopulateNext(t *testing.T) {

	Convey("TestPopulateNext", t, func() {
		Convey("用例1", func() {
			//指针默认值是nil
			tmp := &treeLinkNode{}
			tmp.val = 1
			tmp.left = &treeLinkNode{}
			tmp.left.val = 2
			tmp.right = &treeLinkNode{}
			tmp.right.val = 3

			tmp.left.left = &treeLinkNode{}
			tmp.left.left.val = 4
			tmp.left.right = &treeLinkNode{}
			tmp.left.right.val = 5

			tmp.right.left = &treeLinkNode{}
			tmp.right.left.val = 6
			tmp.right.right = &treeLinkNode{}
			tmp.right.right.val = 7
			PopulateNext(tmp)
			fmt.Printf("%d", tmp.val)

		})

	})

}
