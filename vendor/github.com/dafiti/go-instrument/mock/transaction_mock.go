package mock

import (
	newrelic "github.com/newrelic/go-agent"
	"net/http"
)

type (
	Transaction struct {
		http.ResponseWriter
	}
)

func (t *Transaction) End() error {
	return nil
}

func (t *Transaction) Ignore() error {
	return nil
}

func (t *Transaction) SetName(name string) error {
	return nil
}

func (t *Transaction) NoticeError(err error) error {
	return nil
}

func (t *Transaction) AddAttribute(key string, value interface{}) error {
	return nil
}

func (t *Transaction) StartSegmentNow() newrelic.SegmentStartTime {
	return newrelic.SegmentStartTime{}
}
