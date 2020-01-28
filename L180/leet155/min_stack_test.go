package leet155

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMinStack(t *testing.T) {

	Convey("TestMinStack", t, func() {
		Convey("1", func() {
			stack := NewStack()
			stack.push(-2)
			stack.push(0)
			stack.push(-3)

			if stack.getMin() != -3 {
				t.Errorf("failed")
			}
			stack.pop()

			if stack.top() != 0 {
				t.Errorf("failed")
			}

			if stack.getMin() != -2 {
				t.Errorf("failed")
			}

		})

	})

}
