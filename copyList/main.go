package main

import (
	"fmt"
	"strings"
)

func main() {
	n1 := &Node{
		Val:    1,
		Next:   nil,
		Random: nil,
	}
	n2 := &Node{
		Val:    2,
		Next:   n1,
		Random: n1,
	}
	n3 := &Node{
		Val:    3,
		Next:   n2,
		Random: n1,
	}
	n4 := &Node{
		Val:    4,
		Next:   n3,
		Random: n2,
	}
	n5 := &Node{
		Val:    5,
		Next:   n4,
		Random: n2,
	}
	fmt.Println(n5)

	res := copyRandomList(n5)
	fmt.Println(res)
	fmt.Println(n5)
}

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func (node *Node) String() string {

	res := strings.Builder{}
	for node != nil {
		rv := 0
		if node.Random != nil {
			rv = node.Random.Val
		}
		res.WriteString(fmt.Sprintf("->[%v,%v]", node.Val, rv))
		node = node.Next
	}
	return res.String()
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	//新节点放在旧节点next上
	tmp := head
	for tmp != nil {
		tmp.Next = &Node{
			Val:  tmp.Val,
			Next: tmp.Next,
		}
		tmp = tmp.Next.Next
	}
	//fmt.Println(head)

	//遍历一遍，把random调正确
	pre, now := head, head.Next
	for {
		if pre.Random != nil {
			now.Random = pre.Random.Next
		}

		if now.Next == nil {
			break
		}
		pre = now.Next
		now = pre.Next
	}

	//两个链表分开
	newList := head.Next
	pre, now = head, head.Next
	for {
		pre.Next = now.Next
		if now.Next == nil {
			break
		}
		now.Next = now.Next.Next
		pre = pre.Next
		now = now.Next
	}
	return newList

}
