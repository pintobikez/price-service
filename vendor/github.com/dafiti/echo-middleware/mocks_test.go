package middleware

import (
	nr "github.com/newrelic/go-agent"
	"net/http"
	"time"
)

type (
	Application struct{}

	Transaction struct {
		http.ResponseWriter
	}
)

func (app *Application) StartTransaction(name string, w http.ResponseWriter, r *http.Request) nr.Transaction {
	return new(Transaction)
}

func (app *Application) RecordCustomEvent(eventType string, params map[string]interface{}) error {
	return nil
}

func (app *Application) Shutdown(timeout time.Duration) {
}

func (app *Application) WaitForConnection(timeout time.Duration) error {
	return nil
}

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

func (t *Transaction) StartSegmentNow() nr.SegmentStartTime {
	return nr.SegmentStartTime{}
}
