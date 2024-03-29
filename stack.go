package main
import "fmt"

type Stack struct{
	items[] int
}

func (s *Stack) Push(i int) {
	s.items = append(s.items, i)
}

func (s *Stack) Pop() int {
	l := len(s.items)-1
	toRemove := s.items[l]
	s.items = s.items[:l]

	return toRemove
}

func main(){
	myStack := Stack{}
	fmt.Println(myStack)
	myStack.Push(30)
	myStack.Push(33)
	myStack.Push(39)
	fmt.Println(myStack)
	myStack.Pop()
	fmt.Println(myStack)

	popo := myStack.Pop()
	fmt.Println(popo)

}