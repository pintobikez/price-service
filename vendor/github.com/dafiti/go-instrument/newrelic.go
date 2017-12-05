package instrument

import (
	newrelic "github.com/newrelic/go-agent"
)

type (
	// Structure used to storage the NewRelic transaction
	NewRelic struct {
		txn newrelic.Transaction
	}

	// Structure used to storage the NewRelic External Segment
	NewRelicExternalSegment struct {
		externalSegment newrelic.ExternalSegment
	}

	// Structure used to storage the NewRelic Segment
	NewRelicSegment struct {
		segment newrelic.Segment
	}
)

// Create a external segment
func (nr *NewRelic) ExternalSegment(url string) Segment {
	return NewRelicExternalSegment{
		newrelic.ExternalSegment{
			StartTime: newrelic.StartSegmentNow(nr.txn),
			URL:       url,
		},
	}
}

// Create a segment
func (nr *NewRelic) Segment(name string) Segment {
	return NewRelicSegment{
		newrelic.StartSegment(nr.txn, name),
	}
}

// Ends the external segment
func (es NewRelicExternalSegment) End() error {
	return es.externalSegment.End()
}

// Ends the segment
func (s NewRelicSegment) End() error {
	return s.segment.End()
}

// Set the NewRelic transaction
func (nr *NewRelic) SetTransaction(txn interface{}) {
	nr.txn = txn.(newrelic.Transaction)
}
