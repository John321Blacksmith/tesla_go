package domain

// category object from dataset
type Category struct {
	Label    string   `json:"label"`
	Patterns []string `json:"patterns"`
}

// A sentence object
type Sentence struct {
	Category    string
	LiteralData []string
}
