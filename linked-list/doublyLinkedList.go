package linkedlist

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

// append element on the end on singly linked list
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

// // push element on the beginning
// func (dll *doublyLinkedList) InsertAtTail(data any)

// // Adds a node at a specific position.
// func (dll *doublyLinkedList) InsertAt(index int, data any)

// // Inserts new data right after a specific existing value.
// func (dll *doublyLinkedList) InsertAfter(targetData any, newData any) error

// // Inserts new data right after a specific existing value.
// func (dll *doublyLinkedList) InsertBefore(targetData any, newData any) error

// /*
// 	Deletation ------------------------------------------------------------
// */

// // delete first matched element and return the deleted element
// func (dll *doublyLinkedList) Delete(data any) (bool, any)

// // delete head node.
// func (dll *doublyLinkedList) DeleteHead() (bool, any)

// // delete tail node.
// func (dll *doublyLinkedList) DeleteTail() (bool, any)

// // Removes a node based on its numerical position.
// func (dll *doublyLinkedList) DeleteAt(index int)

// // Keeps the first $n$ elements and deletes the rest.
// func (dll *doublyLinkedList) Truncate(n int)

// /*
// 	Access & Search Methods ------------------------------------------------------------
// */

// // show the head node value
// func (dll *doublyLinkedList) GetHead() (any, error)

// // show the tail node value
// func (dll *doublyLinkedList) GetTail() (any, error)

// // get an element of an given index and a bool status that the index exist or not
// func (dll *doublyLinkedList) GetAt(index int) (bool , any)

// // search an element on linked list and return boolean
// func (dll *doublyLinkedList) Search(data any)  *Node

// // Returns a simple true/false if the value is in the list.
// func (dll *doublyLinkedList) Contains(data any) bool

// /*
// 	Transformation Methods ------------------------------------------------------------
// */

// // Replaces a specific value with a new one.
// func (dll *doublyLinkedList) Update(data, replace any) (bool, any)

// // reverse the linked list
// func (dll *doublyLinkedList) Reverse()

// // sort the linked list
// func (dll *doublyLinkedList) Sort()

// // Scans the list and removes nodes with repeating values
// func (dll *doublyLinkedList) RemoveDuplicates()

// // covert the linked list into slice
// func (dll *doublyLinkedList) ToSlice() []any

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

// // Print the single linked list
// func (dll *doublyLinkedList) PrintList()

// // clear the whole linked list
// func (dll *doublyLinkedList) Clear()

// /*
// 	private helper methods --------------------------------------------------------------------
// */

func (dll *doublyLinkedList) incrementCounter() {
	dll.length++ // just increate by one
}
