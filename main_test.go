package main

import (
	"bufio"
	"bytes"
	"testing"
)

func TestMerge(t *testing.T) {
	var bb bytes.Buffer

	out := bufio.NewWriter(&bb)
	if err := merge("testdata", out, "Jan  2 15:04:05"); err != nil {
		t.Fatalf("Unexpected error: %s", err)
	}

	expected := `Aug  1 00:13:58 Preparation is finished. Wish us good luck. 
Sep  8 00:13:58 We have plenty of food and drinking water.
Sep  19 00:30:24 Thanks to a strong wind we quickly move tovards new adventures.
Sep  25 11:50:04 Did I tell that the weather is just great? It feels more like a vacation.
Sep  30 06:30:41 Today I saw the last seagul, looks like we are far from big land now.
Oct  2 14:00:42 Captain and several sailors are injured. Doctor is here, hopefully we will be able to continue.
Oct  12 04:05:42 Storm was just incredible. And we lost both anchors.
Oct  20 18:44:08 Today is my birthday, and I am afraid that it is my last birthday.
`

	actual := bb.String()

	if actual != expected {
		t.Errorf("\nGot:\n%s\nExpected:\n%s\n", actual, expected)
	}
}
