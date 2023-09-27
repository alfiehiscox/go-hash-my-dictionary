# Exploration of Hashing and Hash Tables

## Goals
- Implement a Hash Table from scratch
- Learning about hash functions and go maps
- Use said Hash Table in a project like scenario

## Project 

*Canoicalisation* project - Given a list of dictionary words *D*, and a set of letters *S* find all the words in *D* that can be made out of letters in *S*. Use a canonical form for each word and a hash table in order to make this search efficient. 

With the above goals in mind, write a CLI game where the user can give sets of letters and will receive the words that can be made by them. 

## Strategy

*These are my initial thoughts before jumping into coding.*

'words.txt' contains all of the words we are going to search. We're going to have a simple interactive CLI, that on startup creates a hash table of the dictionary words. We need to take each word and canonicalise it by creating a sorted set of the characters from it. This sorted set (as a string) is used as the key into the hashmap and the word itself is the value. 

```
Canonical       Maps       Values
aekl             ->        lake, kale, leak 
```

This reduces the complexity of search significantly because all sub-words of *S* will be sorted into the same bucket as the others. We than have to search the bucket for all possible words with the sorted set key and return their values.

Note that this hash-table implementation is slightly different to most because it can store objects with the same key. This may be counterproductive in general, but I think will work okay in this scenario. I could do something similar with a map[string][]string in go, where you don't have to search the whole of the bucket for all occurances of the key. 

## Results

Some notes:
- I've opted for a chained hashtable with doubly linked lists as the buckets. 
- In the source there are three variations on the project. You can comment and uncomment the lines in main.go to see the differences. 
- I haven't added dynamically resizable arrays to my hashtable implementation, although I think this would be pretty easy with slices. 
- As an opionion I think the 'go-way' as outlined below is the best, as will probably become evident. 

I've given three implementations of this idea for retreaving words from a given set of characters.
1. Go implementation: Uses map[string][]string to map the canonical form to each of the actual words. We append words to the slice if they have the same canonical form as the key.
2. HashTable implementation: Tries to emulate the above with a naive implementation of a generic chained hash-table with []string as the value type. On each insert we have to search the indexed bucket to see if the key already exists and update the slice if it does. Otherwise we add to the head of the bucket. Initialisation of the HashTable is a mixture of Search and Insert and is stalled at O(m) time where m is the length of the bucket indexed into. The performance scales with the initial bucket amount chosen. The higher the bucket amount the less 'm' is and the faster Search can perform. 
3. FrequencyTable implementation: I don't know if this is actually what it's called or not, but it's what I described in the 'strategy' section. With this implementation we allow pairs with the same keys to be inserted. Basically all keys are added to the head of the indexed bucket in constant time O(1), meaning initialisation is also constant. Instead of 'Search'ing for the the singular pair with a given key, we search for all the pairs for a given key in the indexed bucket which will always take O(m) where m is the size of the bucket indexed into. This is the same as the worst case in 2.  

The results are as expected when looking at the complexities, this is running the app locally on my measly macbook. The word list is 466550 words long, and I've initialisd the HashTable/FreqTable to have 1000 buckets:

| Implementation | Initialisation | Search            |
| :------------: | :------------: | :---------------: |
| HashTable      |    ~4.5s       |  ~0.03ms - ~0.5ms |
| FreqTable      |    ~0.5s       |  ~0.5ms           |
| Go             |    ~0.5s       |  ~0.002ms         |

So the FreqTable is doing much better than the HashTable on initialisation (as should be expected with constant time), however it should be noted that as the number of buckets chosen approaches n (where n is the number of words) then O(m) approaches constant time 0(1). You can also see that perhaps my hash function is not the best at distributing evenly from the Search times being so mixed. 

And obviously Go is way ahead. What kind of magic is Go doing to replicate the same functionality whilst improving both metrics. Here are the points where I think Go is doing so much better:
- The go hashing function is probably faster and is certainly more evenly distributed then my naive implementation.
- Go has fixed sized buckets of 8 pairs each with a potential pointer to an overflow bucket. This gives faster iteration on search. 
- Go has hash information stored on each bucket to quickly check if the key is in the bucket in constant time, potentially avoiding unnecessary O(m) iteration. 
- Go's bucket array is dynamically sized at a specific load threshold, meaning that we can keep down those search costs at runtime and rely on constant time hash lookups into the bucket array instead of using the overflows. 

## Usage

Prerequiste: Go Installed

1. Clone the repo
2. Run main from root `go run main.go`
3. Type in a sequence of characters
4. See the words made up of the set of these characters

## References 
- The Algorithm Design Manual (Steven S. Skiena)
- [English Words](https://github.com/dwyl/english-words)
- [Inside Go's Map Implementation](https://www.youtube.com/watch?v=Tl7mi9QmLns&t=1376s&ab_channel=GopherAcademy)
