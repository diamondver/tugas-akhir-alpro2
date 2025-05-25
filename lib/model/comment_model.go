package model

// Comment represents a user entity in the system.
// It contains basic identification and authentication information.
type Comment struct {
	// Id is the unique identifier for the comment.
	Id int `json:"id"`

	// UserId is the unique identifier for the user who made the comment.
	UserId int `json:"user_id"`

	// Komentar is the text content of the comment.
	Komentar string `json:"komentar"`

	// Kategori is the category or topic of the comment.
	Kategori string `json:"kategori"`
}
