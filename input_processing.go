package main

// aggregate a list of sentence
// objects
func gather_sentences(fractions []string) []Sentence {
	var s_objects []Sentence = []Sentence{}
	if len(fractions) != 0 {
		for i := 0; i < len(fractions); i++ {
			s_objects = append(
				s_objects,
				Sentence{
					fractions[i],
					"unknown",
					split_to_chunks(fractions[i], " "),
				},
			)
		}
	}
	return s_objects
}

// receive the bytes data
// from the external client
// and split to sentences
func split_to_chunks(data string, sym string) []string {
	chunks := []string{}
	if data != "" {
		flag := 0
		for i := 0; i < len(data); i++ {
			if string(data[i]) == sym {
				chunks = append(chunks, data[flag:i])
				flag = (i + 1)
			}
		}
	}
	return chunks
}

// take the input from the
// the external client
// and process it,
// collecting a list of
// formatted sentences
func ProcessInput(inp string) []Sentence {
	string_list := split_to_chunks(inp, ".")
	sentences := gather_sentences(string_list)
	return sentences
}
