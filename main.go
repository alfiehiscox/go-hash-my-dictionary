package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/alfiehiscox/go-hash-my-dictionary/hashtable"
)

const (
	HT_SIZE = 1000
	PROMT   = ">> "
)

func main() {
	fmt.Println("Welcome to go-hash-my-dictionary.")
	fmt.Println("Type in a continuous string of characters to find like words:")
	start(os.Stdin, os.Stdout)
}

func start(in io.Reader, out io.Writer) {
	ht := DictToFreqHashTable() // Inserting duplicates at head of list
	// ht := DictToHashTable()  // Performing update to []string. Trying to mimic go.
	// m := GoDictToHashTable() // Go implementation

	scanner := bufio.NewScanner(in)
	for {
		fmt.Fprint(out, PROMT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		if strings.Contains(line, " ") {
			fmt.Fprintf(out, "no spaces allowed")
			continue
		}

		h := hash(line)
		before := time.Now()
		results, ok := ht.SearchAll(h) // Search through bucket returning all selected
		// results, ok := ht.Search(h) // Performing update to []string. Trying to mimic go.
		// results, ok := m[h]         // Go implementation

		delta := time.Duration(time.Since(before))
		fmt.Printf("[Took %fs to search HashTable]\n", delta.Seconds())

		if !ok {
			// results is zero value for slice, which is nil
			fmt.Fprintf(out, "no results found")
			continue
		}

		for _, s := range results {
			fmt.Fprintf(out, s+"\n")
		}
	}
}

// Trying to mimic go. Single key per value. Updating value if found.
func DictToHashTable() hashtable.HashTable[[]string] {
	before := time.Now()
	ht := hashtable.New[[]string](HT_SIZE)

	file, err := os.Open("words.txt")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lc := strings.ToLower(scanner.Text())
		h := hash(lc)
		v, _ := ht.Search(h)
		ht.Insert(h, append(v, lc))
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	delta := time.Duration(time.Since(before))
	fmt.Printf("[Took %fs to initialise HashTable]\n", delta.Seconds())
	return ht
}

// Adding everything to hashtable in constant time, with duplicates.
func DictToFreqHashTable() hashtable.HashTable[string] {
	before := time.Now()
	ht := hashtable.New[string](HT_SIZE)

	file, err := os.Open("words.txt")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lc := strings.ToLower(scanner.Text())
		h := hash(lc)
		ht.InsertAll(h, lc)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	delta := time.Duration(time.Since(before))
	fmt.Printf("[Took %fs to initialise HashTable]\n", delta.Seconds())
	return ht
}

// Go implementation
func GoDictToHashTable() map[string][]string {
	before := time.Now()

	m := make(map[string][]string)

	file, err := os.Open("words.txt")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lc := strings.ToLower(scanner.Text())
		h := hash(lc)
		m[h] = append(m[h], lc)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	delta := time.Duration(time.Since(before))
	fmt.Printf("[Took %fs to initialise HashTable]\n", delta.Seconds())

	return m
}

// We assume ascii values, although it shouldn't matter.
func hash(input string) string {
	var runes []rune

	// Make a set of the input
	s := map[rune]bool{}
	for _, r := range input {
		if _, ok := s[r]; !ok {
			s[r] = true
			runes = append(runes, r)
		}
	}

	// Sort by rune
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})

	return string(runes)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
