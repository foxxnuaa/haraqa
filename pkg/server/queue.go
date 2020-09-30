package server

import (
	"io"
	"net/http"

	"github.com/haraqa/haraqa/internal/headers"

	"github.com/haraqa/haraqa/internal/filequeue"
)

//go:generate mockgen -source queue.go -package server -destination queue_mock_test.go
//go:generate goimports -w queue_mock_test.go

var _ Queue = &filequeue.FileQueue{}

// Queue is the interface used by the server to produce and consume messages from different distinct categories called topics
type Queue interface {
	RootDir() string
	Close() error

	ListTopics(prefix, suffix, regex string) ([]string, error)
	CreateTopic(topic string) error
	DeleteTopic(topic string) error
	ModifyTopic(topic string, request headers.ModifyRequest) (*headers.TopicInfo, error)

	Produce(topic string, msgSizes []int64, timestamp uint64, r io.Reader) error
	Consume(group, topic string, id int64, limit int64, w http.ResponseWriter) (int, error)
	SetConsumerOffset(group, topic string, id int64) error
}
