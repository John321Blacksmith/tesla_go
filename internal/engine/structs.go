package engine

// category object from dataset
type Category struct {
	Label    string   `json:"label"`
	Patterns []string `json:"patterns"`
}

// A sentence object
type Sentence struct {
	Data         string
	Category     string
	Literal_data []string
}
