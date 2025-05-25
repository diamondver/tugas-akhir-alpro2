package model

// User represents a user entity in the system.
// It contains basic identification and authentication information.
type User struct {
	// Id is the unique identifier for the user.
	Id int `json:"id"`

	// Username is the unique name used by the user to log in.
	Username string `json:"username"`

	// Password is the user's authentication credential.
	// Note: In a production system, this should be stored as a hash, not plaintext.
	Password string `json:"password"`
}
