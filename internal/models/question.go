package models

import "fmt"

type Question struct {
	Id     uint64
	UserId uint64
	Text   string
}

func (q *Question) String() string {
	return fmt.Sprintf("Question{Id: %v, UserId: %v, Text: \"%v\"}", q.Id, q.UserId, q.Text)
}
