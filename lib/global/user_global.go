package global

import "tugas-besar/lib/model"

// Users is an in-memory storage array that holds up to 255 user records.
// It serves as the persistent storage mechanism for the userRepository implementation.
var Users [255]model.User

// Comments is an in-memory storage array that holds up to 255 comment records.
// It serves as the persistent storage mechanism for the commentRepository implementation.
var Comments [255]model.Comment

// UserCount tracks the current number of users stored in the Users array.
// It's used both as an index for adding new users and for iteration limits when searching.
var UserCount int

// CommentCount tracks the current number of comments stored in the Comments array.
// It's used both as an index for adding new comments and for iteration limits when displaying or processing comments.
var CommentCount int

// IdUserIncrement is a counter used to generate unique IDs for user records.
// It increments each time a new user is created, ensuring each user has a unique identifier.
var IdUserIncrement int

// IdCommentIncrement is a counter used to generate unique IDs for comment records.
// It increments each time a new comment is created, ensuring each comment has a unique identifier.
var IdCommentIncrement int
