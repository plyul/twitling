package main

import (
	"bufio"
	"encoding/hex"
	"log"
	"os"
	"twitling/generated/model"
)

var state = model.State{ //nolint
	Users: map[int64]*model.User{
		1: {Name: "Alice", Following: []int64{2, 4, 5}},
		2: {Name: "Bob", Following: []int64{1}},
		3: {Name: "Carlos", Following: []int64{2, 5}},
		4: {Name: "Dave", Following: nil},
		5: {Name: "Eve", Following: []int64{1, 2, 3, 4}},
	},
	Posts: map[int64]*model.PostDTO{
		1: {Text: "всем чмоки-чмоки", ReplyTo: 0, ToUser: 0, PublishAt: nil},
		2: {Text: "и вам не хворать", ReplyTo: 1, ToUser: 0, PublishAt: nil},
		3: {Text: "привет", ReplyTo: 1, ToUser: 0, PublishAt: nil},
		4: {Text: "сам привет", ReplyTo: 3, ToUser: 0, PublishAt: nil},
		5: {Text: "го в кс?", ReplyTo: 0, ToUser: 0, PublishAt: nil},
		6: {Text: "я пас", ReplyTo: 5, ToUser: 0, PublishAt: nil},
		7: {Text: "а чо так?", ReplyTo: 6, ToUser: 0, PublishAt: nil},
		8: {Text: "нельзя жалеть...", ReplyTo: 7, ToUser: 0, PublishAt: nil},
		9: {Text: "понял, держись там", ReplyTo: 8, ToUser: 0, PublishAt: nil},
	},
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
//		var m model.Message
		hexMsg := scanner.Bytes()
		msg := make([]byte, hex.DecodedLen(len(hexMsg)))
		_, err := hex.Decode(msg, hexMsg)
		if err != nil {
			log.Fatalf("[FATAL]  %v", err)
		}

//		err = proto.Unmarshal(msg, &m)
		if err != nil {
			log.Fatalf("[FATAL]  %v", err)
		}
//		fmt.Printf("Account=%s, Text=%s\n", getUserName(m.Id), m.Text)
	}

	if scanner.Err() != nil {
		log.Fatalf("[FATAL]  %v", scanner.Err())
	}
}

//func getUserName(userId int64) string {
//	for _, u := range state.Users {
//		if u.GetId() == userId {
//			return u.GetName()
//		}
//	}
//	return "Anonymous"
//}
