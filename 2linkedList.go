package main
import "fmt"

type node struct {
	data int
	next *node
}

type linkedList struct {
	length int
	head *node
}

func (l *linkedList) prepend (n *node) {
	second := l.head
	l.head = n
	l.head.next = second
	l.length++
}

func main(){
	myList := linkedList{}
	node1 := &node{data: 99}
	node2 := &node{data: 89}
	node3 := &node{data: 79}

	myList.prepend(node1)
	myList.prepend(node2)
	myList.prepend(node3)

	fmt.Println(myList)
}