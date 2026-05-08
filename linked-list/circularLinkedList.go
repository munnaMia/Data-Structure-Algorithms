package linkedlist

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

// // append element on the end on singly linked list
// func (cll *circularLinkedList) InsertAtHead(data any)

// // push element on the beginning
// func (cll *circularLinkedList) InsertAtTail(data any)

// // Adds a node at a specific position.
// func (cll *circularLinkedList) InsertAt(index int, data any) error

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

// // show the head node value
// func (cll *circularLinkedList) GetHead() (any, error)

// // show the tail node value
// func (cll *circularLinkedList) GetTail() (any, error)

// // get an element of an given index and a bool status that the index exist or not
// func (cll *circularLinkedList) GetAt(index int) (bool, any)

// // search an element on linked list and return boolean
// func (cll *circularLinkedList) Search(data any) bool

// // Returns a simple true/false if the value is in the list.
// func (cll *circularLinkedList) Contains(data any) bool

// /*
// 	Transformation Methods ------------------------------------------------------------
// */

// // Replaces a specific value with a new one.
// func (cll *circularLinkedList) Update(data, replace any) (bool, any)

// // reverse the linked list
// func (cll *circularLinkedList) Reverse()

// // sort the linked list
// func (cll *circularLinkedList) Sort()

// // Scans the list and removes nodes with repeating values
// func (cll *circularLinkedList) RemoveDuplicates()

// // covert the linked list into slice
// func (cll *circularLinkedList) ToSlice() []any

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

// // Print the single linked list
// func (cll *circularLinkedList) PrintList()

// // clear the whole linked list
// func (cll *circularLinkedList) Clear()

// /*
// 	private helper methods --------------------------------------------------------------------
// */

func (cll *circularLinkedList) incrementCounter() {
	cll.length++ // just increate by one
}