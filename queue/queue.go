package queue

import (
	"fmt"

	linkedlist "github.com/munnaMia/Data-Structure-Algorithms/linked-list"
)

type Queue struct {
	doubleLL *linkedlist.DoublyLinkedList
}

func NewQueue() *Queue {
	return &Queue{
		doubleLL: linkedlist.NewDoublyLinkedList(),
	}
}

// Insert an element into the queue
func (q *Queue) Enqueue(data any) any {
	q.doubleLL.InsertAtTail(data)
	newData, _ := q.doubleLL.GetTailData()
	return newData
}

// Remove an element from the queue
func (q *Queue) Dequeue() any {
	if q.IsEmpty() {
		fmt.Println("Queue is empty")
		return nil
	}

	_, deleteData := q.doubleLL.DeleteHead()
	return deleteData
}

// Return the peek element of queue
func (q *Queue) Peek() any {
	if q.IsEmpty() {
		fmt.Println("Queue is empty")
		return nil
	}

	data, _ := q.doubleLL.GetHeadData()
	return data
}

// Print the queue
func (q *Queue) Print() {
	q.doubleLL.PrintList()
}

// Check the queue is empty or not
func (q *Queue) IsEmpty() bool {
	return q.doubleLL.IsEmpty()
}

// Remove all element from the queue
func (q *Queue) Clear() {
	q.doubleLL.Clear()
}

// Show stack is empty or not
func (st *Queue) Size() int {
	return st.doubleLL.Length()
}
