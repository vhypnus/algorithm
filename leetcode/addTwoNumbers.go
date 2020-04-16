/*


*/

package leetcode


type ListNode struct {
	Val int

	Next *ListNode
}

func ArrayToList(array []int) *ListNode {

	var head,tail *ListNode

	for _,val := range array {
		var node = &ListNode{Val:val,Next:nil}

		if head == nil {
			head = node
		}

		if tail != nil {
			tail.Next = node
		} 
		
		tail = node
	


	}

	return head 

}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	var sum int
	var head,tail *ListNode
	for  {
		if l1 == nil && l2 == nil {
			if tail.Val >= 10 {
				tail.Val -= 10
				tail.Next = &ListNode{Val:1,Next:nil}
			}
			break
		}

		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		var node = &ListNode{Val:sum,Next:nil}

		if head == nil {
			head = node
			tail = node
		} else {
			if tail.Val >= 10 {
				node.Val += 1
				tail.Val -= 10
			}

			tail.Next = node

			//reset
			tail = node

		}
		


		//reset 
		sum = 0

	}


	return head
}