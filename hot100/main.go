package main

func main() {
	
}

/*
电话号码的字母组合
https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number/
 */
func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	letterMap := make(map[byte]string)
	letterMap['2'] = "abc"
	letterMap['3'] = "def"
	letterMap['4'] = "ghi"
	letterMap['5'] = "jkl"
	letterMap['6'] = "mno"
	letterMap['7'] = "pqrs"
	letterMap['8'] = "tuv"
	letterMap['9'] = "wxyz"

	var res []string
	db := []byte(digits)
	var letters []string
	for _, d := range db {
		letters = append(letters, letterMap[d])
	}

	var backtrack func([]byte, int)
	backtrack = func(track []byte, start int) {
		if len(track) == len(digits) {
			tmp := make([]byte, len(digits))
			copy(tmp, track)
			res = append(res, string(tmp))
			return
		}

		for i := start; i < len(letters); i++ {
			for j := 0; j < len(letters[i]); j++ {
				track = append(track, letters[i][j])
				backtrack(track, i+1)
				track = track[:len(track)-1]
			}
		}
	}
	backtrack([]byte{}, 0)
	return res
}

type ListNode struct {
	Val int
	Next *ListNode
}

/*
https://leetcode-cn.com/problems/intersection-of-two-linked-lists/
相交链表
给你两个单链表的头节点 headA 和 headB ，请你找出并返回两个单链表相交的起始节点。如果两个链表没有交点，返回 null
让长链表先走一段短链表的长度，然后一起走，看是否相等
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	a, b := headA, headB
	lenA, lenB := 0, 0
	for a != nil && b != nil {
		a = a.Next
		b = b.Next
		lenA ++
		lenB ++
	}
	for a != nil {
		a = a.Next
		lenA ++

	}
	for b != nil {
		b = b.Next
		lenB ++
	}

	if lenA > lenB {
		a = headA
		b = headB
	} else {
		a = headB
		b = headA
		lenA, lenB = lenB, lenA
	}
	for i := 0; i < lenA - lenB; i++ {
		a = a.Next
	}
	for a != b && a != nil && b != nil {
		a = a.Next
		b = b.Next
	}
	return a
}

func countBits(n int) []int {
	res := make([]int, n+1)
	for i := 0; i <= n; i++ {
		count := 0
		num := i
		for num != 0 {
			count += num & 1
			num >>= 1
		}
		res[i] = count
	}
	return res
}