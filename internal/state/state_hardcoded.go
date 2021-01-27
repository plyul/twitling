package state

import "twitling/generated/model"

var state = model.State{ //nolint // unused, убрать, когда станет used
	Users: map[int64]*model.User{
		1: {Name: "Alice", Following: []int64{2, 4, 5}},
		2: {Name: "Bob", Following: []int64{1}},
		3: {Name: "Carlos", Following: []int64{2, 5}},
		4: {Name: "Dave", Following: nil},
		5: {Name: "Eve", Following: []int64{1, 2, 3, 4}},
	},
	Posts: map[int64]*model.Post{
		1: {Text: "hi all", ReplyTo: 0, ToUser: 0, PublishAt: nil},
		2: {Text: "hi to you too", ReplyTo: 1, ToUser: 0, PublishAt: nil},
		3: {Text: "guten tag", ReplyTo: 1, ToUser: 0, PublishAt: nil},
		4: {Text: "me no understdood u", ReplyTo: 3, ToUser: 0, PublishAt: nil},
		5: {Text: "go CS?", ReplyTo: 0, ToUser: 0, PublishAt: nil},
		6: {Text: "pass", ReplyTo: 5, ToUser: 0, PublishAt: nil},
		7: {Text: "why?", ReplyTo: 6, ToUser: 0, PublishAt: nil},
		8: {Text: "not allowed to feel sorry...", ReplyTo: 7, ToUser: 0, PublishAt: nil},
		9: {Text: "got it, hold on there", ReplyTo: 8, ToUser: 0, PublishAt: nil},
	},
}
