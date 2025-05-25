package model

// User represents a user entity in the system.
// It contains basic identification and authentication information.
type Comment struct {
	// komentar is the text content of the comment.
	komentar string `json:"komentar"`

	// kategori is the category or topic of the comment.
	kategori string `json:"kategori"`
}
