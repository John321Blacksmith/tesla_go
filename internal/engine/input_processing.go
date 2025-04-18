package engine

// aggregate a list of sentence
// objects
func gather_sentences(refined_sentences []string) []Sentence {
	var s_objects []Sentence = []Sentence{}
	if len(refined_sentences) != 0 {
		for i := range len(refined_sentences) {
			s_objects = append(
				s_objects,
				Sentence{
					refined_sentences[i],
					"unknown",
					Split_to_chunks(refined_sentences[i], " "),
				},
			)
		}
	}
	return s_objects
}

// check each symbol
// if it's an alphabetic
func check_symbol(sym byte) bool {
	return (sym >= 65 && sym <= 90) ||
		(sym >= 97 && sym <= 122)
}

// check each word for
// having a non alphabetic
// symbol
func refine_sentences(chunks []string) []string {
	var refined_sentences []string
	if len(chunks) != 0 {
		for i := range len(chunks) {
			if len(chunks[i]) != 0 {
				signs_count := 0
				for j := range len(chunks[i]) {
					if !check_symbol(chunks[i][j]) {
						signs_count += 1
					}
				}
				if signs_count == 0 {
					refined_sentences = append(refined_sentences, chunks[i])
				}
			}
		}
	}
	return refined_sentences
}

func get_refined_sentences(data string) []string {
	chunks := Split_to_chunks(data, ".")
	refined_sentences := refine_sentences(chunks)
	return refined_sentences
}

// receive the data
// from the external client
// and split to sentences
func Split_to_chunks(data string, sym string) []string {
	chunks := []string{}
	if data != "" {
		flag := 0
		for i := range len(data) {
			if string(data[i]) == sym {
				chunks = append(chunks, data[flag:i])
				flag = (i + 1)
			}
		}
		chunks = append(chunks, data[flag:])
	}
	return chunks
}

// take the input from the
// the external client
// and process it,
// collecting a list of
// formatted sentences
func ProcessInput(inp string) []Sentence {
	refined_sentences := get_refined_sentences(inp)
	sentences := gather_sentences(refined_sentences)
	return sentences
}
