package basic

import (
	"hello-proto/protogen/basic"
	"log"
	"os"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func BasicWriteUserContentV1() {
	uc := basic.UserContent{
		UserContentId: 90,
		AuthorId:      99,
		Title:         "Protobuf",
		Slug:          "this-is-v1",
		HtmlContent:   "<h1>Protobuf</h1>",
	}

	b, _ := proto.Marshal(&uc)

	if err := os.WriteFile("./public/user_content_v1.bin", b, 0644); err != nil {
		log.Fatalln(err)
	}
	log.Println("Success write:", "user_content_v1.bin")
}

func BasicReadUserContentV1() {
	log.Println("Read v1")
	var uc basic.UserContent

	b, _ := os.ReadFile("./public/user_content_v1.bin")

	if err := proto.Unmarshal(b, &uc); err != nil {
		log.Fatalln(err)
	}

	jsonBytes, _ := protojson.Marshal(&uc)
	log.Println(string(jsonBytes))
}
