package main

import (
	"fmt"
	"hello-proto/basic"
	"log"
	"time"
)

type logWriter struct{}

func (w logWriter) Write(bytes []byte) (int, error) {
	return fmt.Print(time.Now().Format("15:04:05") + "|" + string(bytes))
}

func main() {
	log.SetFlags(0)
	log.SetOutput(new(logWriter))

	// basic.BasicUserGroup()
	// jobsearch.JobSearchSoftware()
	// jobsearch.JobSearchCandidate()
	basic.BasicUser()
	// basic.BasicUnmarshalAnyKnown()
	// basic.BasicUnmarshalAnyNotKnown()
	// basic.BasicUnmarshalAnyIs()
	// basic.BasicOneOf()
	// basic.BasicMap()
	// basic.BasicWriteProto()
	// basic.BasicReadProto()
	// basic.BasicWriteUserContentV1()
	// basic.BasicReadUserContentV1()
}
