package jobsearch

import (
	"hello-proto/protogen/basic"
	"hello-proto/protogen/dummy"
	"hello-proto/protogen/jobsearch"
	"log"

	"google.golang.org/protobuf/encoding/protojson"
)

func JobSearchSoftware() {
	js := jobsearch.JobSoftware{
		JobSoftwareId: 888,
		Application: &basic.Application{
			Version:   "1.0.2.1",
			Name:      "Hello Proto",
			Platforms: []string{"windows", "unix"},
		},
	}

	jsonBytes, _ := protojson.Marshal(&js)
	log.Println("JobSearchSoftware", string(jsonBytes))
}

func JobSearchCandidate() {
	jc := jobsearch.JobCandidate{
		JobCandidateId: 890,
		Application: &dummy.Application{
			ApplicationId:       768,
			ApplicationFullName: "Febriansyah",
			Phone:               "09128391231",
			Email:               "febrian@mail.com",
		},
	}

	jsonBytes, _ := protojson.Marshal(&jc)
	log.Println("JobSearchCandidate", string(jsonBytes))
}
