package basic

import (
	"hello-proto/protogen/basic"
	"log"
	"math/rand"

	"google.golang.org/genproto/googleapis/type/date"
	"google.golang.org/genproto/googleapis/type/latlng"
	"google.golang.org/protobuf/reflect/protoreflect"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func BasicUser() {
	comm := randomCommunicationChannel()
	u := basic.User{
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
			Coordinate: &basic.Address_Coordinate{
				Lat:  103.2391123,
				Long: -203.90221,
			},
		},
		CommunicationChannel: &comm,
		LastLoginTimestamp:   timestamppb.Now(),
		BirthDate:            &date.Date{Year: 1991, Month: 2, Day: 15},
		LastKnownLocation:    &latlng.LatLng{Latitude: 102.203, Longitude: -201.233},
	}

	// proto to json
	jsonBytes, _ := protojson.Marshal(&u)
	jsonString := string(jsonBytes)

	log.Println("Proto to Json: ", jsonString)

	var newUser basic.User
	protojson.Unmarshal(jsonBytes, &newUser)

	log.Println("Json to Proto", &newUser)
}

func randomCommunicationChannel() anypb.Any {

	paperMail := basic.PaperMail{
		PaperMailAddress: "Some paper mail address",
	}

	socialMedia := basic.SocialMedia{
		SocialMediaPlatform: "somePlatform",
		SocialMediaUsername: "user.platf",
	}

	instantMessaging := basic.InstantMessaging{
		InstantMessagingProduct:  "whatsup",
		InstantMessagingUsername: "user.name",
	}

	var a anypb.Any

	switch r := rand.Intn(10) % 3; r {
	case 0:
		anypb.MarshalFrom(&a, &paperMail, proto.MarshalOptions{})
	case 1:
		anypb.MarshalFrom(&a, &socialMedia, proto.MarshalOptions{})
	default:
		anypb.MarshalFrom(&a, &instantMessaging, proto.MarshalOptions{})
	}

	return a
}

func BasicUnmarshalAnyKnown() {
	socialMedia := basic.SocialMedia{
		SocialMediaPlatform: "socialmedia-platform",
		SocialMediaUsername: "socialmedia-username",
	}

	var a anypb.Any
	anypb.MarshalFrom(&a, &socialMedia, proto.MarshalOptions{})

	// knwon tupe (Social Media)
	var sm basic.SocialMedia

	if err := a.UnmarshalTo(&sm); err != nil {
		log.Fatalln(err)
		return
	}

	jsonBytes, _ := protojson.Marshal(&sm)
	log.Println(string(jsonBytes))
}

func BasicUnmarshalAnyNotKnown() {
	comm := randomCommunicationChannel()

	var unmarshaled protoreflect.ProtoMessage

	unmarshaled, err := comm.UnmarshalNew()
	if err != nil {
		log.Fatalln(err)
		return
	}

	log.Println("Type", unmarshaled.ProtoReflect().Descriptor().Name())

	jsonBytes, _ := protojson.Marshal(unmarshaled)
	log.Println(string(jsonBytes))
}

func BasicUnmarshalAnyIs() {
	a := randomCommunicationChannel()
	pm := basic.PaperMail{}

	if a.MessageIs(&pm) {
		if err := a.UnmarshalTo(&pm); err != nil {
			log.Fatalln(err)
		}

		jsonBytes, _ := protojson.Marshal(&pm)
		log.Println(string(jsonBytes))
	} else {
		log.Println("Not PaperMail, but:", a.TypeUrl)
	}

}

func BasicOneOf() {
	socialMedia := basic.SocialMedia{
		SocialMediaPlatform: "socialmedia-platform",
		SocialMediaUsername: "socialmedia-username",
	}

	ecomm := basic.User_SocialMedia{
		SocialMedia: &socialMedia,
	}

	u := basic.User{
		Id:                    919,
		Username:              "Febrian",
		IsActive:              true,
		ElectronicCommChannel: &ecomm,
	}

	jsonBytes, _ := protojson.Marshal(&u)
	jsonString := string(jsonBytes)

	log.Println(jsonString)
}

func BasicMap() {
	sr := map[string]uint32{"fireball": 2, "fly": 1}

	u := basic.User{
		Id:          919,
		Username:    "Febrian",
		IsActive:    true,
		SkillRating: sr,
	}

	jsonBytes, _ := protojson.Marshal(&u)
	jsonString := string(jsonBytes)

	log.Println(jsonString)
}
