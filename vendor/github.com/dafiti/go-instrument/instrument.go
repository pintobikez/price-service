package instrument

type (
	// Instrumentation interface for metrics
	Instrument interface {
		// Added new external segment metric (ex. http clients)
		ExternalSegment(url string) Segment

		// Added a segment metric
		Segment(name string) Segment

		// Define a transaction used in segments
		SetTransaction(txn interface{})
	}

	// Segment interface
	Segment interface {
		// Collection ends for segment
		End() error
	}
)
