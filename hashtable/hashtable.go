// package hashtable holds a 'chaining' hashtable and associated functions.
package hashtable

import (
	"fmt"
	"strings"
)

// Node is the individual elements of the LinkedList
type node[T comparable] struct {
	key  string
	data T
	next *node[T]
	prev *node[T]
}

func (n *node[T]) String() string {
	builder := strings.Builder{}
	builder.WriteString("{key=" + n.key)
	builder.WriteString(",")
	builder.WriteString("value=" + fmt.Sprintf("%v", n.data))
	builder.WriteString("},")
	return builder.String()
}

// Doubly LinkedList has value of type 'T' and key of type 'string'
type linkedList[T comparable] struct {
	head *node[T]
	size int
}

// insert at the head of the list. O(1) complexity.
func (ll *linkedList[T]) insert(key string, data T) {
	// fmt.Println("insert(" + key + ")")
	tmp := &node[T]{key: key, data: data, next: nil, prev: nil}
	if ll.head == nil {
		ll.head = tmp
	} else {
		tmp2 := ll.head
		tmp2.prev = tmp
		ll.head = tmp
		tmp.next = tmp2
	}
	ll.size++
}

// search for string 'key' and got value back. O(n) complexity.
func (ll *linkedList[T]) search(key string) (T, bool) {
	// fmt.Println("search(" + key + ")")
	tmp := ll.head
	for tmp != nil {
		if tmp.key == key {
			return tmp.data, true
		}
		tmp = tmp.next
	}
	return *new(T), false
}

// search for all nodes with the same 'key' but different values.
// returns empty list if none.
func (ll *linkedList[T]) searchAll(key string) []T {
	var xs []T
	tmp := ll.head
	for tmp != nil {
		if tmp.key == key {
			xs = append(xs, tmp.data)
		}
		tmp = tmp.next
	}
	return xs
}

// delete node with given key. Return if successful. O(n) complexity
func (ll *linkedList[T]) delete(key string) bool {
	// fmt.Println("delete(" + key + ")")
	tmp := ll.head
	for tmp != nil {
		if tmp.key == key {
			if tmp.prev == nil {
				// Head of list
				ll.head = tmp.next
				ll.size--
				return true
			} else {
				// Point prev.next to tmp.next
				tmp.prev.next = tmp.next
				ll.size--
				return true
			}
		}
		tmp = tmp.next
	}
	return false
}

func (ll *linkedList[T]) String() string {
	builder := strings.Builder{}
	tmp := ll.head
	builder.WriteString("linkedList{size=" + fmt.Sprint(ll.size) + ", elems=")
	for tmp != nil {
		builder.WriteString(tmp.String())
		tmp = tmp.next
	}
	builder.WriteString("}")
	return builder.String()
}

// HashTable[T] is an array of linkedList[T]
type HashTable[T comparable] []linkedList[T]

// New creates a HashTable of size n, meaning the there are n buckets, not n possible slots.
func New[T comparable](n int) HashTable[T] {
	ht := make(HashTable[T], n)
	return ht
}

// Search(key) returns the value associated with the key.
// Returns an comma-ok value is successful or not.
func (ht HashTable[T]) Search(key string) (T, bool) {
	hash := hash(key, len(ht))
	ll := ht[hash]
	return ll.search(key)
}

func (ht HashTable[T]) SearchAll(key string) []T {
	hash := hash(key, len(ht))
	ll := ht[hash]
	return ll.searchAll(key)
}

// Insert(key, value) adds a pair to the HashTable.
func (ht HashTable[T]) Insert(key string, value T) {
	hash := hash(key, len(ht))
	ht[hash].insert(key, value)
}

// Delete(key) removes a pair from the HashTable.
// Retuns a bool is successful or not.
func (ht HashTable[T]) Delete(key string) bool {
	hash := hash(key, len(ht))
	return ht[hash].delete(key)
}

// String() output's a string for the whole structure.
func (ht HashTable[T]) String() string {
	builder := strings.Builder{}
	builder.WriteString("HashTable[")
	for _, bucket := range ht {
		builder.WriteString(bucket.String())
		builder.WriteString(",")
	}
	builder.WriteString("]")
	return builder.String()
}

func (ht HashTable[T]) AverageBucketLength() float64 {
	var sum float64
	for _, ll := range ht {
		sum += float64(ll.size)
	}
	return sum / float64(len(ht))
}

func (ht HashTable[T]) Size() int {
	var sum int
	for _, ll := range ht {
		sum += ll.size
	}
	return sum
}

// Simple hashing function with assumed ASCCI characters.
func hash(key string, size int) int {
	ba := []byte(key)
	sum := 0
	for _, b := range ba {
		sum += int(b) // Sum ASCCI codes
	}
	return sum % size // return modulo on available address (i.e. size)
}
