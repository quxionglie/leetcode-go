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
