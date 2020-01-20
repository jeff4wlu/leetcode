package leet138

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCpyLstRandomPointer(t *testing.T) {

	Convey("TestCpyLstRandomPointer", t, func() {
		Convey("用例1", func() {

			node1 := new(lstNode)
			node1.val = 1
			node2 := new(lstNode)
			node2.val = 2

			node1.next = node2
			node1.random = node2
			node2.random = node2

			got := CpyLstRandomPointer(node1)

			fmt.Println()
			fmt.Println("old:")
			fmt.Printf(" %d, %p, next: %p, random: %p", node1.val, node1, node1.next, node1.random)
			fmt.Printf(" %d, %p, next: %p, random: %p", node2.val, node2, node2.next, node2.random)
			fmt.Println()

			fmt.Println("new:")
			for got != nil {
				fmt.Printf(" %d, %p, next: %p, random: %p", got.val, got, got.next, got.random)
				got = got.next
			}
			fmt.Println()

		})

	})

}
