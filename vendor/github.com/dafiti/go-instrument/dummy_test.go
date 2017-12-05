package instrument

import (
	"testing"
)

func TestDummyExternalSegment(t *testing.T) {
	dummy := new(Dummy)

	dummy.SetTransaction("transaction")

	if dummy.ExternalSegment("http://localhost").End() != nil {
		t.Fatal("External segment throw error")
	}
}

func TestDummySegment(t *testing.T) {
	dummy := new(Dummy)

	dummy.SetTransaction("transaction")

	if dummy.Segment("segment").End() != nil {
		t.Fatal("Segment throw error")
	}
}
