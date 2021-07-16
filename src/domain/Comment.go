package domain

type Comment struct {
	Comment string
	PostId uint
	CommentBy Profile
	Mentions []Profile

}
