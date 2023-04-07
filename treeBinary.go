package main
import "fmt"

var count int

type Node struct{
	Key int
	Jenis string
	Left *Node
	Right *Node
}

func (n *Node) Insert (k int, jwenis string) {
	if k > n.Key {
		if n.Right == nil {
			n.Right = &Node{Key: k, Jenis: jwenis}
		} else {
			n.Right.Insert(k, jwenis)
		}
	} else if k < n.Key {
		if n.Left == nil {
			n.Left = &Node{Key: k, Jenis: jwenis}
		} else {
			n.Left.Insert(k, jwenis)
		}
	}
}

func (n *Node) Search (k int) bool {
	
	count++
	if n == nil{
		return false
	}

	if k > n.Key {
		return n.Right.Search(k)
	} else if k < n.Key {
		return n.Left.Search(k)
	}

	return true
}

func (n *Node) Print() {
	if n != nil {
		n.Left.Print()
		fmt.Println(n.Key, n.Jenis)
		n.Right.Print()
	}
}

func main() {
	tree := &Node{Key: 100, Jenis: "susu"}
	tree.Insert(50, "koku")
	tree.Insert(120, "jimbe")
	tree.Insert(79, "kokita")
	tree.Insert(39, "mama lemon")
	tree.Insert(250, "mama mia")
	tree.Insert(88, "apa coba")

	fmt.Print(tree)
	fmt.Println(tree.Search(39), count)
	tree.Print()
}