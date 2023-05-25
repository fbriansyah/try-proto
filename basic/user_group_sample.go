package basic

import (
	"hello-proto/protogen/basic"
	"log"

	"google.golang.org/protobuf/encoding/protojson"
)

func BasicUserGroup() {
	u1 := basic.User{
		Id:       99,
		Username: "Febrian",
		Password: []byte("suppersecred"),
		IsActive: true,
		Emails:   []string{"fbriansyah@gmail.com"},
		Gender:   basic.Gender_GENDER_MALE,
		Address: &basic.Address{
			Street:  "JL Buntu",
			City:    "Ngalam",
			Country: "ID",
		},
	}

	u2 := basic.User{
		Id:       98,
		Username: "Rian",
		Password: []byte("suppersecred"),
		IsActive: true,
		Emails:   []string{"ryan@gmail.com"},
		Gender:   basic.Gender_GENDER_MALE,
		Address: &basic.Address{
			Street:  "JL Buntu",
			City:    "Ngalam",
			Country: "ID",
			Coordinate: &basic.Address_Coordinate{
				Lat: 103.2391123,
				Long: -203.90221,
			},
		},
	}

	ug := basic.UserGroup{
		GroupId:   99,
		GroupName: "Coba",
		Roles:     []string{"Admin"},
		Users:     []*basic.User{&u1, &u2},
	}

	jsonBytes, _ := protojson.Marshal(&ug)
	log.Println(string(jsonBytes))
}
