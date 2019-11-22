package _018

import (
	"github.com/whaley/advent-of-code/common"
	"sort"
	"strings"
	"time"
)

type StateChange int

const (
	Switch StateChange = iota
	Sleep  StateChange = iota
	Awake  StateChange = iota
)

type Event struct {
	timeStamp time.Time
	text      string
}

type eventsByTimeStamp []Event

func (a eventsByTimeStamp) Len() int {
	return len(a)
}

func (a eventsByTimeStamp) Less(i, j int) bool {
	return a[i].timeStamp.Before(a[j].timeStamp)
}

func (a eventsByTimeStamp) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

const timestampFormat = "2006-01-02 15:04"

func InputToEvent(input string) (Event, error) {
	tsBeginIndex := strings.IndexRune(input, '[') + 1
	tsEndIndex := strings.IndexRune(input, ']')
	timeStamp, err := time.Parse(timestampFormat, input[tsBeginIndex:tsEndIndex])
	if err != nil {
		return Event{}, err
	}
	text := strings.TrimSpace(input[tsEndIndex+1:])
	return Event{timeStamp, text}, nil
}

func SolveDay04Pt1(lines []string) error {
	var events []Event
	for _, line := range lines {
		event, err := InputToEvent(line)
		if err != nil {
			return err
		}
		events = append(events, event)
	}

	sort.Sort(eventsByTimeStamp(events))
	/*timeChart :=*/ createTimeChartFromEvents(events)
	return nil
}

func createTimeChartFromEvents(events []Event) TimeChart {
	//currentGuardId := math.MaxInt32
	//for _, event := range events {
	//	//TODO: implement
	//}
	return nil
}

type GuardId int
type DateYMD struct {
	year  int
	month int
	day   int
}

func NewDateYMD(t time.Time) DateYMD {
	return DateYMD{t.Year(), int(t.Month()), t.Day()}
}

type TimeChart map[GuardId]map[DateYMD][60]bool

func RunDay04() {
	lines := common.DelimitByNewLine(common.ReadFully("static/2018/day04.txt"))
	err := SolveDay04Pt1(lines)
	common.PanicOnError(err)
}
