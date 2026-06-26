package stack

import linkedlist "github.com/munnaMia/Data-Structure-Algorithms/linked-list"

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
func (st *Stack) Pop() {

}

// Print the stack
func (st *Stack) Print(){
	st.singleLL.PrintList()
}

// Return the top element of the stack
func (st *Stack) Peek() {

}

// Remove all element from the stack
func (st *Stack) Clear() {

}

// Show stack is empty or not
func (st *Stack) IsEmpty() {

}

// Show stack is full or not
func (st *Stack) IsFull() {

}
