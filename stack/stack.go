package stack

import (
	"fmt"

	linkedlist "github.com/munnaMia/Data-Structure-Algorithms/linked-list"
)

type Stack struct {
	singleLL *linkedlist.SinglyLinkedList
}

func NewStack() *Stack {
	return &Stack{
		singleLL: linkedlist.NewSinglyLinkedList(),
	}
}

// Insert an element into the stack
func (st *Stack) Push(data any) any {
	st.singleLL.InsertAtTail(data)

	newData, _ := st.singleLL.GetTailData()

	return newData
}

// Remove an element from the stack
func (st *Stack) Pop() any {
	if st.IsEmpty() {
		fmt.Println("Stack is empty.")
		return nil
	}
	_, data := st.singleLL.DeleteTail()
	return data
}

// Print the stack
func (st *Stack) Print() {
	st.singleLL.PrintList()
}

// Return the top element of the stack
func (st *Stack) Peek() any {
	if st.IsEmpty() {
		fmt.Println("Stack is empty.")
		return nil
	}
	
	data, _ := st.singleLL.GetTailData()
	return data
}

// Remove all element from the stack
func (st *Stack) Clear() {
	st.singleLL.Clear()
}

// Show stack is empty or not
func (st *Stack) IsEmpty() bool {
	return st.singleLL.IsEmpty()
}
