package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"google.golang.org/protobuf/proto"
	"log"
	"miniTwitter/gen/model"
	"os"
)

var state = model.State{
	Users: []*model.User{
		{Id: 1, Name: "Alice", Following: []int64{2, 4, 5}},
		{Id: 2, Name: "Bob", Following: []int64{1}},
		{Id: 3, Name: "Carlos", Following: []int64{2, 5}},
		{Id: 4, Name: "Dave", Following: nil},
		{Id: 5, Name: "Eve", Following: []int64{1, 2, 3, 4}},
	},
	Messages: []*model.Message{
		{Id: 1, Text: "всем чмоки-чмоки", ReplyTo: 0, ToUser: 0, PublishAt: nil},
		{Id: 2, Text: "и вам не хворать", ReplyTo: 1, ToUser: 0, PublishAt: nil},
		{Id: 3, Text: "Всем чмоки-чмоки", ReplyTo: 0, ToUser: 0, PublishAt: nil},
		{Id: 4, Text: "Всем чмоки-чмоки", ReplyTo: 0, ToUser: 0, PublishAt: nil},
		{Id: 5, Text: "Всем чмоки-чмоки", ReplyTo: 0, ToUser: 0, PublishAt: nil},
		{Id: 6, Text: "Всем чмоки-чмоки", ReplyTo: 0, ToUser: 0, PublishAt: nil},
		{Id: 7, Text: "Всем чмоки-чмоки", ReplyTo: 0, ToUser: 0, PublishAt: nil},
		{Id: 8, Text: "Всем чмоки-чмоки", ReplyTo: 0, ToUser: 0, PublishAt: nil},
		{Id: 9, Text: "Всем чмоки-чмоки", ReplyTo: 0, ToUser: 0, PublishAt: nil},
	},
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		var m model.Message
		hexMsg := scanner.Bytes()
		msg := make([]byte, hex.DecodedLen(len(hexMsg)))
		_, err := hex.Decode(msg, hexMsg)
		if err != nil {
			log.Fatalf("[FATAL]  %v", err)
		}

		err = proto.Unmarshal(msg, &m)
		if err != nil {
			log.Fatalf("[FATAL]  %v", err)
		}
		fmt.Printf("Account=%s, Text=%s\n", getUserName(m.Id), m.Text)
	}

	if scanner.Err() != nil {
		log.Fatalf("[FATAL]  %v", scanner.Err())
	}
}

func getUserName(userId int64) string {
	for _, u := range state.Users {
		if u.GetId() == userId {
			return u.GetName()
		}
	}
	return "Anonymous"
}
