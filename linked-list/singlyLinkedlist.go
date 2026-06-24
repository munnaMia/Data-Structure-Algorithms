package linkedlist

import "fmt"

type singlyLinkedList struct {
	head   *Node
	tail   *Node
	length int
}

func NewSinglyLinkedList() *singlyLinkedList {
	return &singlyLinkedList{
		head:   nil,
		tail:   nil,
		length: 0,
	}
}

/*
	Insertion ------------------------------------------------------------
*/

// append element on the beginning on singly linked list
func (sll *singlyLinkedList) InsertAtHead(data any) {
	defer sll.incrementCounter()

	newNode := &Node{
		data: data,
		next: sll.head,
	}

	if sll.IsEmpty() {
		sll.head = newNode
		sll.tail = newNode
		return
	}

	sll.head = newNode
}

// push element on the end
func (sll *singlyLinkedList) InsertAtTail(data any) {
	defer sll.incrementCounter()

	if sll.IsEmpty() {
		sll.head = &Node{
			data: data,
		}

		sll.tail = sll.head
		return
	}

	// storing element on the tail
	sll.tail.next = &Node{
		data: data,
	}

	// move the tail to the end
	sll.tail = sll.tail.next
}

// Adds a node at a specific position. replace data or append empty node.
func (sll *singlyLinkedList) InsertAt(index int, data any) {
	counter := 0

	currentNode := sll.head

	for currentNode != nil {
		if counter == index {
			currentNode.data = data
			return
		}

		currentNode = currentNode.next
		counter++
	}

	for counter <= index {
		sll.tail.next = &Node{}
		sll.tail = sll.tail.next
		sll.incrementCounter()

		if counter == index {
			sll.tail.data = data
			return
		}
		counter++
	}
}

// Inserts new data right after a specific existing value. if nothing match it do nothing just like u.
func (sll *singlyLinkedList) InsertAfter(targetData any, data any) {

	targetNode, _ := sll.Search(targetData)

	if targetNode == nil {
		return
	}

	if sll.tail == targetNode {
		sll.InsertAtTail(data)
		sll.incrementCounter()
		return
	}

	tempNode := targetNode.next
	targetNode.next = &Node{
		data: data,
		next: tempNode,
	}
	sll.incrementCounter()
}

// Inserts new data right after a specific existing value.
func (sll *singlyLinkedList) InsertBefore(targetData any, data any) {
	if sll.IsEmpty() {
		fmt.Println("Linked list is empty.")
		return
	}

	previous := sll.head
	current := sll.head.next

	// if the head match first
	if previous.data == targetData {
		sll.head = &Node{
			data: data,
			next: previous,
		}
		sll.incrementCounter()
		return
	}

	for current != nil {
		if current.data == data {
			previous.next = &Node{
				data: data,
				next: current,
			}
			sll.incrementCounter()
			return
		}

		// move the pointer
		previous = current
		current = current.next
	}

}

// /*
// 	Deletation ------------------------------------------------------------
// */

// delete first matched element and return the deleted element. false mean data not exist or list is empty
func (sll *singlyLinkedList) Delete(data any) (bool, any) {
	if sll.IsEmpty() {
		fmt.Println("Linked list is empty.")
		return false, 0
	}

	previous := sll.head
	current := sll.head.next

	// if the head match first
	if previous.data == data {
		sll.head = sll.head.next
		sll.decrementCounter()
		return true, sll.head.data
	}

	for current != nil {
		if current.data == data {
			oldData := current.data
			previous.next = current.next // unlink the match data
			sll.decrementCounter()
			return true, oldData
		}
		// move the pointer
		previous = current
		current = current.next
	}

	return false, 0
}

// // delete head node.
// func (sll *singlyLinkedList) DeleteHead() (bool, any)

// // delete tail node.
// func (sll *singlyLinkedList) DeleteTail() (bool, any)

// // Removes a node based on its numerical position.
// func (sll *singlyLinkedList) DeleteAt(index int)

// Keeps the first $n$ elements and deletes the rest.
func (sll *singlyLinkedList) Truncate(n int) error {
	if n == 0 {
		return nil
	}

	if sll.Length() < n {
		return fmt.Errorf("not enough elements in the linked list")
	}

	currentNode, err := sll.GetAt(n - 1)
	if err != nil {
		return err
	}

	sll.head = currentNode.next

	return nil
}

// /*
// 	Access & Search Methods ------------------------------------------------------------
// */

// show the head node value
func (sll *singlyLinkedList) GetHead() (*Node, error) {
	if sll.IsEmpty() {
		return nil, fmt.Errorf("Linked list is empty.")
	}
	return sll.head, nil
}

// show the tail node value
func (sll *singlyLinkedList) GetTail() (*Node, error) {
	if sll.IsEmpty() {
		return nil, fmt.Errorf("Linked list is empty.")
	}
	return sll.tail, nil
}

// get an node by index
func (sll *singlyLinkedList) GetAt(index int) (*Node, error) {
	if sll.IsEmpty() {
		fmt.Println("Linked list is empty.")
		return nil, fmt.Errorf("linked list is empty")
	}

	currentNode := sll.head
	counter := 0

	for currentNode != nil && counter < sll.length {
		if counter == index {
			return currentNode, nil
		}
		currentNode = currentNode.next
		counter++
	}

	return nil, fmt.Errorf("linked list is empty")
}

// // search an element on linked list and return boolean
func (sll *singlyLinkedList) Search(data any) (*Node, error) {
	if sll.IsEmpty() {
		return nil, fmt.Errorf("linked list is empty")
	}

	currentNode := sll.head

	for currentNode != nil {
		if currentNode.data == data {
			return currentNode, nil
		}
		currentNode = currentNode.next
	}

	return nil, fmt.Errorf("element not found")
}

// Returns a simple true/false if the value is in the list.
func (sll *singlyLinkedList) Contains(data any) bool {
	_, err := sll.Search(data)

	if err != nil {
		return false
	}

	return true
}

/*
	Transformation Methods ------------------------------------------------------------
*/

// Replaces a specific value with a new one.
func (sll *singlyLinkedList) Update(data, replace any) (*Node, error) {
	if sll.IsEmpty() {
		return nil, fmt.Errorf("Linked list is empty.")
	}

	targetNode, err := sll.Search(data)

	if err != nil {
		return nil, err
	}

	oldData := targetNode

	targetNode.data = replace

	return oldData, nil
}

// reverse the linked list
func (sll *singlyLinkedList) Reverse() {
	if sll.IsEmpty() {
		fmt.Println("Linked list is empty.")
		return
	}

	tempLinkedList := NewSinglyLinkedList()

	currentNode := sll.head

	for currentNode != nil {
		tempLinkedList.InsertAtHead(currentNode.data)
		currentNode = currentNode.next
	}

	sll.head = tempLinkedList.head
}

// sort the linked list
// func (sll *singlyLinkedList) Sort()

// Scans the list and removes nodes with repeating values
func (sll *singlyLinkedList) RemoveDuplicates() {
	if sll.head == nil {
		return
	}

	seen := make(map[any]bool)

	current := sll.head
	seen[current.data] = true

	for current.next != nil {
		if seen[current.next.data] {
			current.next = current.next.next
		} else {
			seen[current.next.data] = true
			current = current.next

		}

	}
}

// covert the linked list into slice
func (sll *singlyLinkedList) ToSlice() []any {
	if sll.IsEmpty() {
		fmt.Println("Linked list is empty.")
		return []any{}
	}

	sllSlice := []any{}
	currentNode := sll.head
	for currentNode != nil {
		sllSlice = append(sllSlice, currentNode.data)

		currentNode = currentNode.next
	}

	return sllSlice
}

/*
	Metadata & Utility Methods ------------------------------------------------------------
*/

// tell how many element the linked list have
func (sll *singlyLinkedList) Length() int {
	return sll.length
}

// check the linked list is empty or not
func (sll *singlyLinkedList) IsEmpty() bool {
	return sll.head == nil
}

// Print the single linked list
func (sll *singlyLinkedList) PrintList() {
	if sll.IsEmpty() {
		fmt.Println("Linked list is empty.")
		return
	}

	currentNode := sll.head

	for currentNode != nil {
		fmt.Println("Data :", currentNode.data)
		currentNode = currentNode.next
	}
}

// clear the whole linked list
func (sll *singlyLinkedList) Clear() {
	sll.head = nil
	sll.tail = nil
}

/*
	private helper methods --------------------------------------------------------------------
*/

// increment after eash insertion
func (sll *singlyLinkedList) incrementCounter() {
	sll.length++ // just increate by one
}

// decrement after eash deletation
func (sll *singlyLinkedList) decrementCounter() {
	sll.length-- // just increate by one
}
