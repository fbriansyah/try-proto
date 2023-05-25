package basic

import (
	"hello-proto/protogen/basic"
	"log"
)

func BasicHello() {
	h := basic.User{
		Id:       99,
		Username: "Febrian",
		IsActive: true,
		Password: []byte("Test"),
	}

	log.Println(&h)
}
