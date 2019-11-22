package _018

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
	"time"
)

func TestInputToEvent(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    string
		expected Event
	}{
		{"Shift Begin", "[1518-10-31 23:50] Guard #10 begins shift", Event{time.Date(1518, 10, 31, 23, 50, 0, 0, time.UTC),  "Guard #10 begins shift" }},
		{"Falls Asleep", "[1518-11-01 00:05] falls asleep", Event{time.Date(1518, 11, 1, 0, 5, 0, 0, time.UTC),  "falls asleep"}},
		{"Wakes Up", "[1518-11-01 00:25] wakes up", Event{time.Date(1518, 11, 1, 0, 25, 0, 0, time.UTC),  "wakes up"}},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			event, err := InputToEvent(test.input)
			if assert.NoError(t, err) {
				assert.Equal(t, test.expected, event)
			}

		})
	}
}

func TestSortEventsByTimestamp(t *testing.T) {
	events := []Event{
		{time.Date(1518, 11, 1, 0, 5, 0, 0, time.UTC),  "should be second" },
		{time.Date(1518, 11, 1, 0, 15, 0, 0, time.UTC),  "should be fourth" },
		{time.Date(1518, 11, 1, 0, 0, 0, 0, time.UTC),  "should be first" },
		{time.Date(1518, 11, 1, 0, 10, 0, 0, time.UTC),  "should be third" },
	}
	expected := []Event{
		{time.Date(1518, 11, 1, 0, 0, 0, 0, time.UTC),  "should be first" },
		{time.Date(1518, 11, 1, 0, 5, 0, 0, time.UTC),  "should be second" },
		{time.Date(1518, 11, 1, 0, 10, 0, 0, time.UTC),  "should be third" },
		{time.Date(1518, 11, 1, 0, 15, 0, 0, time.UTC),  "should be fourth" },
	}
	sort.Sort(eventsByTimeStamp(events))
	assert.Equal(t, expected, events)
}

