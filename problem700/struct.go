package problem700

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type ListNode struct {
	Val  int
	Next *ListNode
}

type Node struct {
	Val      int
	Children []*Node
}

func isLinkListEq(n1, n2 *ListNode) bool {
	for n1 != nil || n2 != nil {
		if (n1 == nil && n2 != nil) || (n1 != nil && n2 == nil) {
			return false
		}
		if n1.Val != n2.Val {
			return false
		}
		n1 = n1.Next
		n2 = n2.Next
	}
	return true
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
