package domain

type Comment struct {
	Id string
	Comment string
	PostId uint
	CommentBy Profile
	Mentions []Profile

}
