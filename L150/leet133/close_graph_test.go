package leet133

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCloseGraph(t *testing.T) {

	Convey("TestCloseGraph", t, func() {
		Convey("用例1", func() {

			node1 := new(graphNode)
			node1.val = 1
			node2 := new(graphNode)
			node2.val = 2
			node3 := new(graphNode)
			node3.val = 3
			node4 := new(graphNode)
			node4.val = 4
			node1.neighbors = append(node1.neighbors, node2)
			node1.neighbors = append(node1.neighbors, node4)
			node2.neighbors = append(node2.neighbors, node1)
			node2.neighbors = append(node2.neighbors, node3)
			node3.neighbors = append(node3.neighbors, node2)
			node3.neighbors = append(node3.neighbors, node4)
			node4.neighbors = append(node4.neighbors, node1)
			node4.neighbors = append(node4.neighbors, node3)

			dict := make(map[int]*graphNode)

			got := CloseGraph(node1, &dict)

			fmt.Println()
			fmt.Println("old:")
			fmt.Printf(" %d, %p", node1.val, node1)
			fmt.Printf(" %d, %p", node2.val, node2)
			fmt.Printf(" %d, %p", node3.val, node3)
			fmt.Printf(" %d, %p", node4.val, node4)
			fmt.Println()

			fmt.Println("new:")
			for _, v := range dict {
				fmt.Printf(" %d, %p", v.val, v)
			}
			fmt.Println()

			fmt.Printf("%d", got.val)

		})

	})

}
