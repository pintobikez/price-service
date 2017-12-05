package instrument

import (
	"github.com/dafiti/go-instrument/mock"
	"testing"
)

func TestNewRelicExternalSegment(t *testing.T) {
	nr := new(NewRelic)

	nr.SetTransaction(&mock.Transaction{})

	if nr.ExternalSegment("http://localhost").End() != nil {
		t.Fatal("External segment throw error")
	}
}

func TestNewRelicSegment(t *testing.T) {
	nr := new(NewRelic)

	nr.SetTransaction(&mock.Transaction{})

	if nr.Segment("segment").End() != nil {
		t.Fatal("Segment throw error")
	}
}
