package problem100

import (
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	newNode := mergeTwoLists(newListNode([]int{1, 2, 4}), newListNode([]int{1, 3, 4}))
	assertList(t, newNode, []int{1, 1, 2, 3, 4, 4})
	assertList(t, mergeTwoLists(newListNode([]int{}), newListNode([]int{})), []int{})
	assertList(t, mergeTwoLists(newListNode([]int{0}), newListNode([]int{})), []int{0})
}
