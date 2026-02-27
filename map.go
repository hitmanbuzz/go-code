//go:build ignore

package main

import "fmt"

const MAP_SIZE = 100

type HashTable[T any] struct {
	key   string
	value T
}

type Node[T any] struct {
	data HashTable[T]
	next *Node[T]
}

type HashMap[T any] struct {
	arr  [MAP_SIZE]*Node[T]
	size int
}

func NewMap[T any]() *HashMap[T] {
	var items [MAP_SIZE]*Node[T]
	return &HashMap[T]{
		arr:  items,
		size: 0,
	}
}

func (m *HashMap[T]) Insert(k string, v T) {
	idx := ComputeHash(k) % MAP_SIZE
	node := &Node[T]{
		data: HashTable[T]{
			key:   k,
			value: v,
		},
		next: m.arr[idx],
	}

	m.arr[idx] = node
	m.size++
}

func (m *HashMap[T]) Get(k string) *T {
	idx := ComputeHash(k) % MAP_SIZE
	curr := m.arr[idx]

	for curr != nil {
		table := curr.data
		if table.key == k {
			return &table.value
		}

		curr = curr.next
	}

	return nil
}

func (m *HashMap[T]) Delete(k string) bool {
	isDelete := false
	idx := ComputeHash(k) % MAP_SIZE
	curr := m.arr[idx]

	if curr == nil {
		return false
	}

	if curr.data.key == k {
		m.arr[idx] = curr.next
		m.size--
		return true
	}

	for curr.next != nil {
		table := curr.next.data
		if table.key == k {
			// found the key, now remove it and shift all the later nodes towards it
			curr.next = curr.next.next
			m.size--
			return true
		}
		curr = curr.next
	}

	return isDelete
}

// modulo the hash result with the size of the array
func ComputeHash(input string) uint {
	const p uint = 31
	const m uint = 1e9 + 9
	var hashValue uint = 0
	var pPow uint = 1
	for _, c := range input {
		hashValue = (hashValue + (uint(c)-uint('a')+1)*pPow) % m
		pPow = (pPow * p) % m
	}

	return hashValue
}

func main() {
	hashMap := NewMap[int]()
	hashMap.Insert("John", 18)
	hashMap.Insert("Abraham", 24)
	hashMap.Insert("Jack", 30)
	hashMap.Insert("Luis", 50)
	hashMap.Insert("Maple", 25)
	hashMap.Insert("Ron", 20)

	john := hashMap.Get("John")
	if john == nil {
		fmt.Println("Key: `John` not found")
	} else {
		fmt.Printf("For John: %d\n", *john)
	}

	abra := hashMap.Get("Abraham")
	if abra == nil {
		fmt.Println("Key: `Abraham` not found")
	} else {
		fmt.Printf("For Abraham: %d\n", *abra)
	}

	luis := hashMap.Get("Luis")
	if luis == nil {
		fmt.Println("Key: `Luis` not found")
	} else {
		fmt.Printf("For Luis: %d\n", *luis)
	}

	maple := hashMap.Get("Maple")
	if maple == nil {
		fmt.Println("Key: `Maple` not found")
	} else {
		fmt.Printf("For Maple: %d\n", *maple)
	}

	rohn := hashMap.Get("Apple")
	if rohn == nil {
		fmt.Println("Key: `Apple` not found")
	} else {
		fmt.Printf("For Apple: %d\n", *rohn)
	}

	status := hashMap.Delete("Jack")
	if !status {
		fmt.Println("Key not found, so unable to delete it")
	} else {
		// delete so not found in the table
		jack := hashMap.Get("Jack")
		if jack == nil {
			fmt.Println("Key: `Jack` not found")
		} else {
			fmt.Printf("For Jack: %d\n", *jack)
		}
	}
}
