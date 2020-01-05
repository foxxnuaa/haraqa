package protocol

import (
	"github.com/pkg/errors"
)

//go:generate protoc --gogofaster_out=plugins=grpc:. grpc.proto

var (
	//ErrTopicExists is returned if a CreateTopic request is made to an existing topic
	ErrTopicExists = errors.New("topic already exists")
	//ErrTopicDoesNotExist is returned if a Request is made on a non existent topic
	ErrTopicDoesNotExist = errors.New("topic does not exist")
	//ErrUndefined is returned in the absence of a known error
	ErrUndefined = errors.New("undefined error occurred")
)

// ErrorToResponse converts a standard error to a 2 byte response for stream responses
func ErrorToResponse(err error) [2]byte {
	if err == nil {
		return [2]byte{0, 0}
	}
	switch errors.Cause(err) {
	case ErrTopicDoesNotExist:
		return [2]byte{0, 1}
	default:
		return [2]byte{255, 255}
	}
}

// ResponseToError converts a 2 byte response to an error
func ResponseToError(resp [2]byte) error {
	switch resp {
	case [2]byte{0, 0}:
		return nil
	case [2]byte{0, 1}:
		return ErrTopicDoesNotExist
	case [2]byte{255, 255}:
		return ErrUndefined
	default:
		return ErrUndefined
	}
}
