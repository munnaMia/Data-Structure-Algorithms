package linkedlist

import (
	"fmt"
)

type doublyLinkedList struct {
	head   *Node
	tail   *Node
	length int
}

func NewDoublyLinkedList() *doublyLinkedList {
	return &doublyLinkedList{
		head:   nil,
		tail:   nil,
		length: 0,
	}
}

// /*
// 	Insertion ------------------------------------------------------------
// */

// append element on the beginning on singly linked list
func (dll *doublyLinkedList) InsertAtHead(data any) {
	defer dll.incrementCounter()

	// create a new node
	newNode := &Node{
		prev: nil,
		data: data,
		next: dll.head,
	}

	if dll.IsEmpty() {
		dll.head = newNode
		dll.tail = newNode
		return
	}

	// previous head node prev pointer set to new node
	dll.head.prev = newNode

	dll.head = newNode
}

// push element on the end
func (dll *doublyLinkedList) InsertAtTail(data any) {
	defer dll.incrementCounter()

	if dll.IsEmpty() {
		dll.head = &Node{
			prev: nil,
			data: data,
			next: nil,
		}
		dll.tail = dll.head
		return
	}

	dll.tail.next = &Node{
		prev: dll.tail,
		data: data,
		next: nil,
	}

	dll.tail = dll.tail.next
}

// Adds a node at a specific position.
func (dll *doublyLinkedList) InsertAt(index int, data any) {
	if dll.IsEmpty() {
		fmt.Println("Linked list is empty.")
		return
	}

	if index == 0 {
		newNode := &Node{
			prev: nil,
			data: data,
			next: dll.head,
		}
		dll.head.prev = newNode
		dll.head = newNode
		dll.incrementCounter()
		return
	}

	counter := 1

	current := dll.head.next // start from index 1

	for current != nil {
		if counter == index {
			newNode := &Node{
				prev: current.prev,
				data: data,
				next: current,
			}
			current.prev.next = newNode
			dll.incrementCounter()
			return
		}

		current = current.next
		counter++
	}

	for counter <= index {
		dll.tail.next = &Node{}
		dll.tail = dll.tail.next
		dll.incrementCounter()

		if counter == index {
			dll.tail.data = data
			return
		}
		counter++
	}
}

// Inserts new data right after a specific existing value.
func (dll *doublyLinkedList) InsertAfter(targetData any, data any) {
	if dll.IsEmpty() {
		fmt.Println("Linked list is empty.")
		return
	}

	targetNode, _ := dll.Search(targetData)

	if targetNode == nil {
		return
	}

	// insert at tail
	if dll.tail == targetNode {
		dll.InsertAtTail(data)
		dll.incrementCounter()
		return
	}

	newNode := &Node{
		prev: targetNode,
		data: data,
		next: targetNode.next,
	}

	targetNode.next.prev = newNode
	targetNode.next = newNode

	dll.incrementCounter()
}

// Inserts new data right before a specific existing value.
func (dll *doublyLinkedList) InsertBefore(targetData any, data any) {
	if dll.IsEmpty() {
		fmt.Println("Linked list is empty.")
		return
	}

	targetNode, _ := dll.Search(targetData)

	if targetNode == nil {
		return
	}

	// insert at head
	if dll.head == targetNode {
		dll.InsertAtHead(data)
		dll.incrementCounter()
		return
	}

	newNode := &Node{
		prev: targetNode.prev,
		data: data,
		next: targetNode,
	}

	targetNode.prev.next = newNode
	targetNode.prev = newNode
	dll.incrementCounter()
}

// /*
// 	Deletation ------------------------------------------------------------
// */

// delete first matched element and return the deleted element
func (dll *doublyLinkedList) Delete(data any) (bool, any) {
	if dll.IsEmpty() {
		fmt.Println("Linked list is empty.")
		return false, 0
	}
	targetNode, _ := dll.Search(data)

	if dll.length == 1 {
		if targetNode.data == dll.head.data {
			oldData := dll.head.data
			dll.head = nil
			dll.tail = nil
			dll.decrementCounter()
			return true, oldData
		}
	}

	if dll.head == targetNode {
		oldData := dll.head.data
		dll.head = dll.head.next
		dll.head.prev = nil
		dll.decrementCounter()
		return true, oldData
	}

	if dll.tail == targetNode {
		oldData := dll.tail.data
		dll.tail = dll.tail.prev
		dll.tail.next = nil
		dll.decrementCounter()
		return true, oldData
	}

	oldData := targetNode.data
	preNode := targetNode.prev
	postNode := targetNode.next

	preNode.next = postNode
	postNode.prev = preNode
	dll.decrementCounter()
	return true, oldData
}

// delete head node.
func (dll *doublyLinkedList) DeleteHead() (bool, any) {
	if dll.IsEmpty() {
		fmt.Println("Linked list is empty.")
		return false, 0
	}
	oldData := dll.head.data

	if dll.length == 1 {
		dll.head = nil
		dll.tail = nil
		dll.decrementCounter()
		return true, oldData
	}

	dll.head = dll.head.next
	dll.head.prev = nil
	dll.decrementCounter()
	return true, oldData
}

// delete tail node.
func (dll *doublyLinkedList) DeleteTail() (bool, any) {
	if dll.IsEmpty() {
		fmt.Println("Linked list is empty.")
		return false, 0
	}
	oldData := dll.tail.data

	if dll.length == 1 {
		dll.head = nil
		dll.tail = nil
		dll.decrementCounter()
		return true, oldData
	}

	dll.tail = dll.tail.prev
	dll.tail.next = nil
	dll.decrementCounter()

	return true, oldData

}

// Removes a node based on its numerical position.
func (dll *doublyLinkedList) DeleteAt(index int) (bool, any) {
	if dll.IsEmpty() {
		fmt.Println("Linked list is empty.")
		return false, 0
	}

	counter := 0

	current := dll.head

	if dll.length == 1 {
		if counter == index {
			oldData := dll.head.data
			dll.head = nil
			dll.tail = nil
			dll.decrementCounter()
			return true, oldData
		}
	}

	if index == 0 {
		oldData := dll.head.data
		dll.head = dll.head.next
		dll.head.prev = nil
		dll.decrementCounter()
		return true, oldData
	}

	if index == dll.length-1 {
		oldData := dll.tail.data
		dll.tail = dll.tail.prev
		dll.tail.next = nil
		dll.decrementCounter()
		return true, oldData
	}

	for counter <= index {
		if counter == index {
			oldData := current.data
			preNode := current.prev
			postNode := current.next

			preNode.next = postNode
			postNode.prev = preNode

			dll.decrementCounter()
			return true, oldData
		}
		counter++
		current = current.next
	}

	return false, nil
}

// Keeps the first $n$ elements and deletes the rest.
func (dll *doublyLinkedList) Truncate(n int) error {
	if n == 0 {
		return nil
	}

	if dll.Length() < n {
		return fmt.Errorf("not enough elements in the linked list")
	}

	currentNode, err := dll.GetAt(n - 1)
	if err != nil {
		return err
	}
	
	dll.head = currentNode.next

	return nil
}

// /*
// 	Access & Search Methods ------------------------------------------------------------
// */

// show the head node value
func (dll *doublyLinkedList) GetHead() (*Node, error) {
	if dll.IsEmpty() {
		return nil, fmt.Errorf("Linked list is empty.")
	}
	return dll.head, nil
}

// show the tail node value
func (dll *doublyLinkedList) GetTail() (*Node, error) {
	if dll.IsEmpty() {
		return nil, fmt.Errorf("Linked list is empty.")
	}
	return dll.tail, nil
}

// get an element of an given index and a bool status that the index exist or not
func (dll *doublyLinkedList) GetAt(index int) (*Node, error) {
	if dll.IsEmpty() {
		fmt.Println("Linked list is empty.")
		return nil, fmt.Errorf("linked list is empty")
	}

	currentNode := dll.head
	counter := 0

	for currentNode != nil && counter < dll.length {
		if counter == index {
			return currentNode, nil
		}
		currentNode = currentNode.next
		counter++
	}

	return nil, fmt.Errorf("linked list is empty")

}

// search an element on linked list and return boolean
func (dll *doublyLinkedList) Search(data any) (*Node, error) {
	if dll.IsEmpty() {
		return nil, fmt.Errorf("linked list is empty")
	}

	currentNode := dll.head

	for currentNode != nil {
		if currentNode.data == data {
			return currentNode, nil
		}
		currentNode = currentNode.next
	}

	return nil, fmt.Errorf("element not found")
}

// Returns a simple true/false if the value is in the list.
func (dll *doublyLinkedList) Contains(data any) bool {
	_, err := dll.Search(data)

	if err != nil {
		return false
	}

	return true
}

// /*
// 	Transformation Methods ------------------------------------------------------------
// */

// Replaces a specific value with a new one.
func (dll *doublyLinkedList) Update(data, replace any) (*Node, error) {
	if dll.IsEmpty() {
		return nil, fmt.Errorf("Linked list is empty.")
	}

	targetNode, err := dll.Search(data)

	if err != nil {
		return nil, err
	}

	oldData := targetNode

	targetNode.data = replace

	return oldData, nil
}

// // reverse the linked list
// func (dll *doublyLinkedList) Reverse()

// // Scans the list and removes nodes with repeating values
// func (dll *doublyLinkedList) RemoveDuplicates()

// covert the linked list into slice
func (dll *doublyLinkedList) ToSlice() []any {
	if dll.IsEmpty() {
		fmt.Println("Linked list is empty.")
		return []any{}
	}

	sllSlice := []any{}
	currentNode := dll.head
	for currentNode != nil {
		sllSlice = append(sllSlice, currentNode.data)

		currentNode = currentNode.next
	}

	return sllSlice
}

// /*
// 	Metadata & Utility Methods ------------------------------------------------------------
// */

// tell how many element the linked list have
func (dll *doublyLinkedList) Length() int {
	return dll.length
}

// check the linked list is empty or not
func (dll *doublyLinkedList) IsEmpty() bool {
	return dll.head == nil
}

// Print the single linked list
func (dll *doublyLinkedList) PrintList() {
	if dll.IsEmpty() {
		fmt.Println("Linked list is empty.")
		return
	}

	currentNode := dll.head

	for currentNode != nil {
		fmt.Println("Data :", currentNode.data)
		currentNode = currentNode.next
	}
}

// clear the whole linked list
func (dll *doublyLinkedList) Clear() {
	dll.head = nil
	dll.tail = nil
}

// /*
// 	private helper methods --------------------------------------------------------------------
// */

// increment after eash deletation
func (dll *doublyLinkedList) incrementCounter() {
	dll.length++ // just increate by one
}

// decrement after eash deletation
func (dll *doublyLinkedList) decrementCounter() {
	dll.length--
}
