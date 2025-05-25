package global

import "tugas-besar/lib/model"

// Users is an in-memory storage array that holds up to 255 user records.
// It serves as the persistent storage mechanism for the userRepository implementation.
var Users [255]model.User
var Comments [255]model.Comment

// UserCount tracks the current number of users stored in the Users array.
// It's used both as an index for adding new users and for iteration limits when searching.
var UserCount int
var CommentCount int
