package domain

type Comment struct {
	Id uint
	Comment string
	PostId uint
	CommentBy Profile
	Mentions []Profile

}
