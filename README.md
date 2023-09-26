# Exploration of Hashing and Hash Tables

## Goals
- Implement a Hash Table from scratch (closed-addressed and open-addressed)
- Learning about hash functions
- Use said Hash Table in a project like scenario

## Project 

*Canoicalisation* project - Given a list of dictionary words *D*, and a set of letters *S* find all the words in *D* that can be made out of letters in *S*. Use a canonical forms for each word and hashing in order to make this search efficient. As a follow up question, which set of *k* letters can be used to make the *most* dictionary words?

With the above goals in mind, write a CLI tool where the user can give sets of letters and will receive the words that can be made by them. Addiontally, given a max command and a *k* value return the set of *k* letters that makes the *most* dictionary words. 

## Strategy

'words.txt' contains all of the words we are going to search. We're going to have a simple CLI, that on startup, we create a hash table of the words. We need to take each word and canonicalise it by creating a sorted set of the characters from it. This sorted set (as a string) is used as the key into the hashmap and the word itself is the value. 

This reduces the complexity of search significantly because all sub-words of *S* will be sorted into the same bucket as the others. We than have to search the bucket for all possible words with the sorted set key and return their values.

Note that my hash-table implementation is slightly different to most because it can store objects with the same key. This may be is counterproductive in general, but it seems to work okay in the scenario. You could do something similar with a map[string][]string in go. 

## Usage

Prerequiste: Go Installed

1. Clone the repo
2. Run main from root `go run main.go`
3. Type in a sequence of characters
4. See the words made up of the set of these characters

## References 
- The Algorithm Design Manual (Steven S. Skiena)
- [English Words](https://github.com/dwyl/english-words)