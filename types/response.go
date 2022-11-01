package types

import "fmt"

type Response struct {
	Ok          bool
	Result      *Result
	Error_code  int
	Description string
}

type Result struct {
	Message_id  int64
	Sender_chat *Sender_chat
	Chat        *Chat
	Date        int64
	Text        string
}

type Sender_chat struct {
	Id    int64
	Title string
	Type  string
}

type Chat struct {
	Id    int64
	Title string
	Type  string
}

func (r *Response) String() string {
	if r == nil {
		return "<nil>"
	}
	return fmt.Sprintf("%+v", *r)
}

func (r *Result) String() string {
	if r == nil {
		return "<nil>"
	}
	return fmt.Sprintf("%+v", *r)
}

func (s *Sender_chat) String() string {
	if s == nil {
		return "<nil>"
	}
	return fmt.Sprintf("%+v", *s)
}

func (c *Chat) String() string {
	if c == nil {
		return "<nil>"
	}
	return fmt.Sprintf("%+v", *c)
}
