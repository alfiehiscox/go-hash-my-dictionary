package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"

	"github.com/alfiehiscox/go-hash-my-dictionary/hashtable"
)

const (
	HT_SIZE = 100
	PROMT   = ">> "
)

func main() {
	fmt.Println("Welcome to go-hash-my-dictionary.")
	fmt.Println("Type in a continuous string of characters to find like words:")
	start(os.Stdin, os.Stdout)
}

func start(in io.Reader, out io.Writer) {
	ht := DictToHashTable()

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
		results := ht.SearchAll(h)

		if len(results) == 0 {
			fmt.Fprintf(out, "no results found")
		}

		for _, s := range results {
			fmt.Fprintf(out, s+"\n")
		}
	}
}

func DictToHashTable() hashtable.HashTable[string] {
	ht := hashtable.New[string](HT_SIZE)

	file, err := os.Open("words.txt")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lc := strings.ToLower(scanner.Text())
		h := hash(lc)
		ht.Insert(h, lc)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return ht
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
