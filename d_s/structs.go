package d_s

// category object from dataset
type Category struct {
	Label    string   `json:"label"`
	Patterns []string `json:"patterns"`
}

// A sentence object
type Sentence struct {
	data         string
	category     string
	literal_data []string
}
