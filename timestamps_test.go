package main

import (
	"testing"
)

func TestFirstEvent(t *testing.T) {
	var (
		blank = ""
		oct   = "Oct  12 04:05:42 Storm was just incredible. And we lost both anchors."
		sep   = "Sep  19 00:30:24 Thanks to a strong wind we quickly move tovards new adventures."
		aug   = "Aug  1 00:13:58 Preparation is finished. Wish us good luck."
	)

	records := []string{oct, blank, sep, aug}
	sorter := newEventSorter(len(records), "Jan  2 15:04:05")

	firstEvent, err := sorter.firstEventIndex(records)
	if err != nil {
		t.Error("Unexpected error:", err)
	}
	actual := records[firstEvent]
	if actual != aug {
		t.Errorf("Got %s as first event, expected: %s", records[firstEvent], actual)
	}
}
