package instrument

type (
	// Dummy Structure
	Dummy struct {
	}

	// Dummy External Structure
	DummyExternalSegment struct {
	}

	// Dummy Segment Structure
	DummySegment struct {
	}
)

// Create a external segment
func (d *Dummy) ExternalSegment(url string) Segment {
	return new(DummyExternalSegment)
}

// Create a segment
func (d *Dummy) Segment(name string) Segment {
	return new(DummySegment)
}

// Ends the external segment
func (es DummyExternalSegment) End() error {
	return nil
}

// Ends the segment
func (s DummySegment) End() error {
	return nil
}

// Set the NewRelic transaction
func (d *Dummy) SetTransaction(txn interface{}) {
}
