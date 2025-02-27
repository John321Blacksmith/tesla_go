package main

import (
	"fmt"
)

// A sentence object
type Sentence struct {
	data         string
	category     string
	literal_data []string
}

// Am aggregation of all sentence
// objects
type SentencesList struct {
	main_context string
	sentences    []Sentence
}

type Recognized struct {
	sentences []Sentence
}

type Unrecognized struct {
	sentences []Sentence
}

// split the sentence data to the
// separate words
func get_literal_data(inp string, sym string) {}

// format each string
// fraction to the Sentence
// object
func format_sentence(item string) Sentence {}

// aggregate a list of sentence
// objects
func gather_sentences(fractions []string) []Sentence {}

// receive the bytes data
// from the external client
// bytes -> string -> []string -> []Sentence
func process_input(inp byte) []string {}

func main() {
	input := "This is a test input. All the data is not relevant. The facts described are fictional."
	fmt.Println(input)

}
