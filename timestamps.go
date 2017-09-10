package main

import (
	"sort"
	"strings"
	"time"
)

type timestamps []time.Time

// Len is the number of elements in the collection.
func (ts timestamps) Len() int { return len(ts) }

// Less reports whether the element with
// index i should sort before the element with index j.
func (ts timestamps) Less(i, j int) bool {
	return ts[i].UnixNano() < ts[j].UnixNano()
}

// Swap swaps the elements with indexes i and j.
func (ts timestamps) Swap(i, j int) {
	ts[i], ts[j] = ts[j], ts[i]
}

type eventSorter struct {
	timeformat string
	ts         []time.Time
	sorted     timestamps
}

func newEventSorter(nReaders int, tf string) *eventSorter {
	return &eventSorter{
		timeformat: tf,
		ts:         make([]time.Time, nReaders),
		sorted:     make(timestamps, nReaders),
	}
}

func (es *eventSorter) firstEventIndex(records []string) (int, error) {
	for i, record := range records {
		if record == "" {
			es.ts[i] = time.Now()
			continue
		}
		timePrefix := strings.TrimSpace(record[:len(es.timeformat)+1])
		t, err := time.Parse(es.timeformat, timePrefix)
		if err != nil {
			return 0, err
		}
		es.ts[i] = t
	}

	copy(es.sorted, es.ts)
	sort.Sort(es.sorted)
	first := es.sorted[0]

	for i, s := range es.ts {
		if s == first {
			return i, nil
		}
	}

	panic("We have lost count of everything.")
}
