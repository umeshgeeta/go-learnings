package lessons

type ListNode struct {
	Val  int
	Next *ListNode
}

func GetIntersectionNode(headA, headB *ListNode) *ListNode {
	lengthA := length(headA)
	lengthB := length(headB)

	var foundIntersection *ListNode = nil
	if lengthA > 0 && lengthB > 0 {
		diff := lengthA - lengthB
		dist := lengthB
		hereA := headA
		hereB := headB
		if diff < 0 {
			diff = diff * -1
			dist = lengthA
			hereB = move(hereB, diff)
		} else {
			hereA = move(hereA, diff)
		}
		j := 0
		for hereA != nil && hereB != nil && j < dist {
			if hereA == hereB {
				foundIntersection = hereA
				break
			}
			hereA = hereA.Next
			hereB = hereB.Next
			j++
		}
	}
	return foundIntersection
}

func move(head *ListNode, howMuch int) *ListNode {
	here := head
	i := 0
	for i < howMuch {
		here = here.Next
		i++
	}
	return here
}

func length(head *ListNode) int {
	length := 0
	here := head
	for here != nil {
		here = here.Next
		length += 1
	}
	return length
}
