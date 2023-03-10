package main

import "fmt"

type Node struct {
    value int
    next  *Node
    prev  *Node
}

type DoublyLinkedList struct {
    head *Node
    tail *Node
}

func (list *DoublyLinkedList) Add(value int) {
    newNode := &Node{value: value}

    if list.head == nil {
        list.head = newNode
        list.tail = newNode
    } else {
        list.tail.next = newNode
        newNode.prev = list.tail
        list.tail = newNode
    }
}

func (list *DoublyLinkedList) PrintForward() {
    current := list.head
    for current != nil {
        fmt.Print(current.value)
        current = current.next
    }
}

func (list *DoublyLinkedList) PrintBackward() {
    current := list.tail
	fmt.Print("\n")
    for current != nil {
        fmt.Print(current.value)
        current = current.prev
    }
}

func main() {
    list := DoublyLinkedList{}
    list.Add(1)
    list.Add(2)
    list.Add(3)
    list.PrintForward()  // Output: 1 2 3
    list.PrintBackward() // Output: 3 2 1
}
