package problem000

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func createLinkList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	var head, tail *ListNode
	for _, num := range nums {
		if head == nil {
			head = &ListNode{
				Val:  num,
				Next: nil,
			}
			tail = head
		} else {
			tail.Next = &ListNode{
				Val:  num,
				Next: nil,
			}
			tail = tail.Next
		}
	}
	return head
}

func isTreeEq(n1, n2 *TreeNode) bool {
	if n1 == nil && n2 == nil {
		return true
	}
	if (n1 != nil && n2 == nil) || (n1 == nil && n2 != nil) {
		return false
	}
	if n1.Val != n2.Val {
		return false
	}
	return isTreeEq(n1.Left, n2.Left) && isTreeEq(n1.Right, n2.Right)
}
