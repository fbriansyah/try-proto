package basic

import (
	"hello-proto/protogen/basic"
	"log"
	"os"

	"google.golang.org/protobuf/proto"

	"google.golang.org/protobuf/encoding/protojson"
)

func BasicWriteProto() {
	sr := map[string]uint32{"fireball": 2, "fly": 1}

	u := basic.User{
		Id:          919,
		Username:    "Febrian",
		IsActive:    true,
		SkillRating: sr,
	}

	uBytes, err := proto.Marshal(&u)
	if err != nil {
		log.Fatalln(err)
	}

	os.WriteFile("./protofile", uBytes, 0644)
}

func BasicReadProto() {
	var u basic.User
	uBytes, err := os.ReadFile("./protofile")
	if err != nil {
		log.Fatalln(err)
	}
	proto.Unmarshal(uBytes, &u)

	jsonBytes, _ := protojson.Marshal(&u)
	jsonString := string(jsonBytes)

	log.Println(jsonString)
}
