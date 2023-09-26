package hashtable

import (
	"testing"
)

func TestHash(t *testing.T) {
	size := 10
	tests := []struct {
		input  string
		output int
	}{
		{"test", 8},
		{"TesTinG", 0},
		{"raNd!", 2},
		{"12rand2", 0},
	}

	for _, tt := range tests {
		hash := hash(tt.input, size)
		if hash != tt.output {
			t.Errorf("hash code did not match. got=%d expected=%d", hash, tt.output)
		}
	}
}

func TestLinkedList(t *testing.T) {
	expected := createTestList()

	ll := linkedList[string]{}
	ll.insert("three", "3")
	ll.insert("two", "2")
	ll.insert("one", "1")

	if ll.size != expected.size {
		t.Fatalf("size incorrect. got=%d wanted=%d", ll.size, expected.size)
	}

	compareLists[string](t, expected, ll)
	_, ok := ll.search("two")
	if !ok {
		t.Fatalf("Could not find key=two")
	}

	l2 := linkedList[string]{}
	l2.insert("three", "3")
	l2.insert("one", "1")

	ll.delete("two")

	if ll.size != l2.size {
		t.Fatalf("size incorrect after delete. got=%d wanted=%d", ll.size, l2.size)
	}

	compareLists[string](t, l2, ll)
}

func createTestList() linkedList[string] {
	node1 := &node[string]{key: "one", data: "1"}
	node2 := &node[string]{key: "two", data: "2"}
	node1.next = node2
	node3 := &node[string]{key: "three", data: "3"}
	node2.next = node3
	node2.prev = node1
	return linkedList[string]{
		head: node1,
		size: 3,
	}
}

func compareLists[T comparable](t *testing.T, expected linkedList[T], actual linkedList[T]) {
	exTmp := expected.head
	acTmp := actual.head

	for exTmp != nil {
		t.Log("Tests")

		if exTmp.key != acTmp.key {
			t.Fatalf("keys don't match. got=%s wanted=%s", exTmp.key, acTmp.key)
		}

		if exTmp.data != acTmp.data {
			t.Fatalf("data doesn't match. got=%v wanted %v", exTmp.data, acTmp.data)
		}

		exTmp = exTmp.next
		acTmp = acTmp.next
	}
}

func TestHashTable(t *testing.T) {
	ht := New[int](10)
	ht.Insert("one", 1)
	ht.Insert("two", 2)
	ht.Insert("three", 3)
	ht.Insert("neo", 69) // Should have the same key as 'one'

	one, ok := ht.Search("one")
	if !ok {
		t.Fatalf("could not find 'one' in %v", ht.String())
	}
	if one != 1 {
		t.Fatalf("wrong value. expected=%d got=%d", 1, one)
	}

	ht.Delete("neo")

	if _, ok := ht.Search("neo"); ok {
		t.Fatalf("found 'neo' after deletion")
	}
}
