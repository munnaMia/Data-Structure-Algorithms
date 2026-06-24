package linkedlist

import "fmt"

type circularLinkedList struct {
	head   *Node
	tail   *Node
	length int
}

func NewCircularLinkedList() *circularLinkedList {
	return &circularLinkedList{
		head:   nil,
		tail:   nil,
		length: 0,
	}
}

// /*
// 	Insertion ------------------------------------------------------------
// */

// append element on the end on singly linked list
func (cll *circularLinkedList) InsertAtHead(data any) {
	defer cll.incrementCounter()

	newNode := &Node{
		prev: cll.tail,
		data: data,
		next: cll.head,
	}

	if cll.IsEmpty() {
		cll.head = newNode
		cll.tail = newNode
	}
	cll.tail.next = newNode
	cll.head.prev = newNode
	cll.head = newNode
}

// push element on the beginning
func (cll *circularLinkedList) InsertAtTail(data any) {
	defer cll.incrementCounter()

	newNode := &Node{
		prev: cll.tail,
		data: data,
		next: cll.head,
	}

	if cll.IsEmpty() {
		cll.head = newNode
		cll.tail = newNode
	}

	cll.tail.next = newNode
	cll.head.prev = newNode
	cll.tail = newNode
}

// // Adds a node at a specific position.
// func (cll *circularLinkedList) InsertAt(index int, data any)

// // Inserts new data right after a specific existing value.
// func (cll *circularLinkedList) InsertAfter(targetData any, newData any) error

// // Inserts new data right after a specific existing value.
// func (cll *circularLinkedList) InsertBefore(targetData any, newData any) error

// /*
// 	Deletation ------------------------------------------------------------
// */

// // delete first matched element and return the deleted element
// func (cll *circularLinkedList) Delete(data any) (bool, any)

// // delete head node.
// func (cll *circularLinkedList) DeleteHead() (bool, any)

// // delete tail node.
// func (cll *circularLinkedList) DeleteTail() (bool, any)

// // Removes a node based on its numerical position.
// func (cll *circularLinkedList) DeleteAt(index int)

// // Keeps the first $n$ elements and deletes the rest.
// func (cll *circularLinkedList) Truncate(n int)

// /*
// 	Access & Search Methods ------------------------------------------------------------
// */

// show the head node value
func (cll *circularLinkedList) GetHead() (any, error) {
	if cll.IsEmpty() {
		return nil, fmt.Errorf("Linked list is empty.")
	}
	return cll.head, nil
}

// show the tail node value
func (cll *circularLinkedList) GetTail() (any, error) {
	if cll.IsEmpty() {
		return nil, fmt.Errorf("Linked list is empty.")
	}
	return cll.tail, nil
}

// // get an element of an given index and a bool status that the index exist or not
// func (cll *circularLinkedList) GetAt(index int) (bool, any)

// // search an element on linked list and return boolean
// func (cll *circularLinkedList) Search(data any)  *Node

// // Returns a simple true/false if the value is in the list.
// func (cll *circularLinkedList) Contains(data any) bool

// /*
// 	Transformation Methods ------------------------------------------------------------
// */

// // Replaces a specific value with a new one.
// func (cll *circularLinkedList) Update(data, replace any) (bool, any)

// // reverse the linked list
// func (cll *circularLinkedList) Reverse()

// // Scans the list and removes nodes with repeating values
// func (cll *circularLinkedList) RemoveDuplicates()

// covert the linked list into slice
func (cll *circularLinkedList) ToSlice() []any {
	if cll.IsEmpty() {
		fmt.Println("Linked list is empty.")
		return nil
	}

	current := cll.head
	slc := make([]any, 0)

	for {
		slc = append(slc, current.data)
		current = current.next

		if current == cll.head {
			break
		}
	}

	return slc
}

// /*
// 	Metadata & Utility Methods ------------------------------------------------------------
// */

// tell how many element the linked list have
func (cll *circularLinkedList) Length() int {
	return cll.length
}

// check the linked list is empty or not
func (cll *circularLinkedList) IsEmpty() bool {
	return cll.head == nil
}

// Print the single linked list
func (cll *circularLinkedList) PrintList() {
	if cll.IsEmpty() {
		fmt.Println("Linked list is empty.")
		return
	}

	current := cll.head

	for {
		fmt.Println("Data :", current.data)
		current = current.next

		if current == cll.head {
			break
		}
	}
}

// clear the whole linked list
func (cll *circularLinkedList) Clear() {
	cll.head = nil
	cll.tail = nil
}

// /*
// 	private helper methods --------------------------------------------------------------------
// */

// increment after eash deletation
func (cll *circularLinkedList) incrementCounter() {
	cll.length++ // just increate by one
}

// decrement after eash deletation
func (cll *circularLinkedList) decrementCounter() {
	cll.length--
}
