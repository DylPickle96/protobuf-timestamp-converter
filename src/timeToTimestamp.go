package main

import (
	"log"
	"time"

	"github.com/golang/protobuf/ptypes"
	//ptypes "github.com/golang/protobuf/ptypes"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
)

// CreateProtobufTimestamp converts a string to a date then to a protobuf timestamp
func CreateProtobufTimestamp(timeString string) *timestamp.Timestamp {
	var timestamp *timestamp.Timestamp

	t, err := time.Parse(time.RFC3339, timeString)

	if err != nil {
		log.Fatalf("Error parsing provided time. Error: %v", err)
	}

	timestamp, err = ptypes.TimestampProto(t)

	if err != nil {
		log.Fatalf("Error converting provided time to a Protobuf Timestamp. Error: %v", err)
	}

	return timestamp
}
