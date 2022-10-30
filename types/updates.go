package types

import "fmt"

type Update struct {
	Update_id    int64
	Message      *Message
	Channel_post *Channel_post
}

type Message struct {
	Message_id int64
	From       struct {
		Id            int64
		Is_bot        bool
		First_name    string
		Language_code string
	}
	Chat struct {
		Id         int64
		First_name string
		Type       string
	}
	Date uint64
	Text string
}

type Channel_post struct {
	Message_id  int64
	Sender_chat struct {
		Id    int64
		Title string
		Type  string
	}
	Chat struct {
		Id    int64
		Title string
		Type  string
	}
	Date uint64
	Text string
}

func (m *Message) String() string {
	if m == nil {
		return "<nil>"
	}
	return fmt.Sprintf("%+v", *m)
}

func (c *Channel_post) String() string {
	if c == nil {
		return "<nil>"
	}
	return fmt.Sprintf("%+v", *c)
}
