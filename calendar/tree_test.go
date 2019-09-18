package calendar_test

import (
	"testing"
	"time"

	"github.com/sajeevany/conflict/calendar"
)

func TestInsert(t *testing.T) {

	//Define layout and time ranges for testing
	layout := "2006-01-02 15:04:05.000000000 +0000 UTC"

	//Time ranges Sept 6 2019 4 AM - 10 AM
	s2019_9_6_4, _ := time.Parse(layout, "2019-09-06 4:00:00.000000000 +0000 UTC")
	s2019_9_6_5, _ := time.Parse(layout, "2019-09-06 5:00:00.000000000 +0000 UTC")
	s2019_9_6_6, _ := time.Parse(layout, "2019-09-06 6:00:00.000000000 +0000 UTC")
	s2019_9_6_8, _ := time.Parse(layout, "2019-09-06 8:00:00.000000000 +0000 UTC")
	s2019_9_6_7, _ := time.Parse(layout, "2019-09-06 7:00:00.000000000 +0000 UTC")
	s2019_9_6_10, _ := time.Parse(layout, "2019-09-06 10:00:00.000000000 +0000 UTC")

	e0, _ := calendar.Event{Start: s2019_9_6_6, End: s2019_9_6_8}.ToEventNode()
	e1, _ := calendar.Event{Start: s2019_9_6_4, End: s2019_9_6_8}.ToEventNode()
	e2, _ := calendar.Event{Start: s2019_9_6_4, End: s2019_9_6_5}.ToEventNode()
	e3, _ := calendar.Event{Start: s2019_9_6_6, End: s2019_9_6_7}.ToEventNode()
	e4, _ := calendar.Event{Start: s2019_9_6_7, End: s2019_9_6_10}.ToEventNode()
	e5, _ := calendar.Event{Start: s2019_9_6_4, End: s2019_9_6_10}.ToEventNode()

	calendar.Insert(&e0, &e1)
	calendar.Insert(&e0, &e2)
	calendar.Insert(&e0, &e3)
	calendar.Insert(&e0, &e4)
	calendar.Insert(&e0, &e5)

}
