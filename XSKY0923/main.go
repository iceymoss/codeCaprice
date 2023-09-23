package main

import (
	"fmt"
)

type Node struct {
	Key  int
	Next *Node
}

type HashTable struct {
	tableSize int
	Table     []*Node
}

func NewHashTable(size int) *HashTable {
	return &HashTable{
		tableSize: size,
		Table:     make([]*Node, size),
	}
}

func (ht *HashTable) Hash(key int) int {
	return key % ht.tableSize
}

func (ht *HashTable) Insert(key int) {
	index := ht.Hash(key)
	node := &Node{Key: key}
	if ht.Table[index] == nil {
		ht.Table[index] = node
	} else {
		current := ht.Table[index]
		for current.Next != nil {
			current = current.Next
		}
		current.Next = node
	}
}

func (ht *HashTable) Search(key int) bool {
	index := ht.Hash(key)
	current := ht.Table[index]
	for current != nil {
		if current.Key == key {
			return true
		}
		current = current.Next
	}
	return false
}

func (ht *HashTable) Delete(key int) {
	index := ht.Hash(key)
	current := ht.Table[index]
	var prev *Node
	for current != nil {
		if current.Key == key {
			if prev != nil {
				prev.Next = current.Next
			} else {
				ht.Table[index] = current.Next
			}
			return
		}
		prev = current
		current = current.Next
	}
}

func main() {
	size := 10
	hashTable := NewHashTable(size)
	keys := []int{0, 10, 20, 30, 40, 50, 60, 70, 80, 90}

	for _, key := range keys {
		hashTable.Insert(key)
	}

	fmt.Println("Search for 10:", hashTable.Search(10)) // Output: true
	fmt.Println("Search for 15:", hashTable.Search(15)) // Output: false

	hashTable.Delete(10)
	fmt.Println("After deleting 10, Search for 10:", hashTable.Search(10)) // Output: false
}
