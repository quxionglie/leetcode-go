package problem100

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
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
