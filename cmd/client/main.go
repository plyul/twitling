package main

import (
	"encoding/hex"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"log"
	"twitling/generated/model"
)

const userId = 1 //nolint

var texts = [...]string{
	"Hello there. Thanks for the follow. Did you notice, that I am an egg? A talking egg? Damn!",
	"Thanks mate! Feel way better now",
	"Yeah that is crazy",
	"Hi",
	"Thanks",
	"Okay",
	"I just wrote 'protobug' instead of 'protobuf' in a report, it's a good summary of Monday morning.",
}

func main() {
	var jl, pl float64
	for account, text := range texts {
		m := model.PostDTO{
			Text: text,
			ReplyTo: 0,
			ToUser: int64(account),
			PublishAt: nil,
		}
		protobufString, err := proto.Marshal(&m)
		if err != nil {
			log.Fatalf("%v", err)
		}
		s := hex.EncodeToString(protobufString)
		jsonString, err := protojson.Marshal(&m)
		if err != nil {
			log.Fatalf("%v", err)
		}
		pl += float64(len(protobufString))
		log.Printf("Message(hex, len=%d)=%s", len(protobufString), s)
		jl += float64(len(jsonString))
		log.Printf("Message(json, len=%d)=%s", len(jsonString), jsonString)
		log.Printf("Protobuf efficiency for message: %.2f", float64(len(protobufString))/float64(len(jsonString)))
	}
	log.Printf("Cumulative protobuf efficiency: %.2f", pl/jl)
}
