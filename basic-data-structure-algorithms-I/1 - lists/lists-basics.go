package main

import (
	"container/list"
	"fmt"
)

func main() {

	// Arrays
	// Arrays are fixed-size sequences of elements of the same type
	var fruitsA [3]string
	fruitsA[0] = "apple"
	fmt.Println(fruitsA)

	fruitsB := [3]string{"apple", "banana", "cherry"}
	fmt.Println(fruitsB)

	fruitsC := [...]string{"apple", "banana", "cherry"}
	fmt.Println(fruitsC)

	// Slices
	// Slice is a reference to an array and manages the length and capacity of the array
	var fruitsD []string
	fruitsD = append(fruitsD, "apple")
	fmt.Println(fruitsD)

	fruitsE := []string{"apple", "banana", "cherry"}
	fmt.Println(fruitsE)

	fruitsF := make([]string, 3) // only length is specified
	fruitsF[0] = "apple"

	fruitsG := make([]string, 3, 5) // length and capacity are specified
	fruitsG[0] = "apple"

	// Lists
	// Lists are arrays with dynamic sizes
	myList := list.New()

	// Push elements to the front and back
	element1 := myList.PushFront(10)
	element2 := myList.PushBack(20)

	// Insert elements after and before specific elements
	myList.InsertAfter(15, element1) // Insert 15 after 10
	myList.InsertBefore(5, element2) // Insert 5 before 20

	// Accessing elements
	fmt.Println("First element:", myList.Front().Value) // Should print 5
	fmt.Println("Last element:", myList.Back().Value)   // Should print 20

	// Length of the list
	fmt.Println("Length of list:", myList.Len()) // Should print 4

	// Remove an element
	myList.Remove(element1)                            // Remove the element containing 10
	fmt.Println("Length after removal:", myList.Len()) // Should print 3

	// Move an element to the front and back
	myList.MoveToFront(element2) // Move 20 to the front
	myList.MoveToBack(element2)  // Move 20 back to the end

	// Initializing the list (clearing it)
	myList.Init()
	fmt.Println("Length after initialization:", myList.Len()) // Should print 0

	// Re-adding elements after initialization
	myList.PushBack(30)
	myList.PushFront(40)

	// Iterating through the list with Next and Prev
	fmt.Println("Elements after re-adding...")
	for e := myList.Front(); e != nil; e = e.Next() {
		fmt.Println("Element:", e.Value)
	}

	// Demonstrating Next and Prev
	fmt.Println("Using Next and Prev:")
	if firstElem := myList.Front(); firstElem != nil {
		fmt.Println("First element:", firstElem.Value) // Should print 40
		if nextElem := firstElem.Next(); nextElem != nil {
			fmt.Println("Next element:", nextElem.Value) // Should print 30
		}
	}

	if lastElem := myList.Back(); lastElem != nil {
		fmt.Println("Last element:", lastElem.Value) // Should print 30
		if prevElem := lastElem.Prev(); prevElem != nil {
			fmt.Println("Previous element:", prevElem.Value) // Should print 40
		}
	}
}
