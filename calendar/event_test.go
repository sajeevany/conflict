package calendar_test

import (
	"testing"
	"time"

	"github.com/sajeevany/conflict/calendar"
)

func TestOverlap(t *testing.T) {

	//Define layout and time ranges for testing
	layout := "2006-01-02 15:04:05.000000000 +0000 UTC"

	//Time ranges Sept 6 2019 4 AM - 10 AM
	s2019_9_6_4, _ := time.Parse(layout, "2019-09-06 4:00:00.000000000 +0000 UTC")
	s2019_9_6_5, _ := time.Parse(layout, "2019-09-06 5:00:00.000000000 +0000 UTC")
	s2019_9_6_6, _ := time.Parse(layout, "2019-09-06 6:00:00.000000000 +0000 UTC")
	s2019_9_6_8, _ := time.Parse(layout, "2019-09-06 8:00:00.000000000 +0000 UTC")
	s2019_9_6_7, _ := time.Parse(layout, "2019-09-06 7:00:00.000000000 +0000 UTC")
	s2019_9_6_10, _ := time.Parse(layout, "2019-09-06 10:00:00.000000000 +0000 UTC")

	nodes := []struct {
		prim calendar.Event
		sec  calendar.Event
		exp  bool
	}{
		//less than
		{calendar.Event{s2019_9_6_4, s2019_9_6_5}, calendar.Event{s2019_9_6_6, s2019_9_6_8}, false},
		//right overlap
		{calendar.Event{s2019_9_6_4, s2019_9_6_7}, calendar.Event{s2019_9_6_6, s2019_9_6_8}, true},
		//left overlap
		{calendar.Event{s2019_9_6_6, s2019_9_6_8}, calendar.Event{s2019_9_6_7, s2019_9_6_10}, true},
		//greater than
		{calendar.Event{s2019_9_6_4, s2019_9_6_5}, calendar.Event{s2019_9_6_7, s2019_9_6_8}, false},
	}

	//Run Overlaps operation for each pair and compare to expected value
	for _, c := range nodes {
		val := c.prim.Overlaps(c.sec)
		if val != c.exp {
			t.Errorf("Expected overlap result of nodes %v and %v to be %v but was %v. Failure.", c.prim, c.sec, c.exp, val)
		}
	}
}

//Test EventNode#CompareTo function
func TestCompareTo(t *testing.T) {

	//Define layout and time ranges for testing
	layout := "2006-01-02 15:04:05.000000000 +0000 UTC"

	//Time ranges Sept 6 2019 4 AM - 10 AM
	s2019_9_6_4, _ := time.Parse(layout, "2019-09-06 4:00:00.000000000 +0000 UTC")
	s2019_9_6_5, _ := time.Parse(layout, "2019-09-06 5:00:00.000000000 +0000 UTC")
	s2019_9_6_6, _ := time.Parse(layout, "2019-09-06 6:00:00.000000000 +0000 UTC")
	s2019_9_6_8, _ := time.Parse(layout, "2019-09-06 8:00:00.000000000 +0000 UTC")
	s2019_9_6_7, _ := time.Parse(layout, "2019-09-06 7:00:00.000000000 +0000 UTC")
	s2019_9_6_9, _ := time.Parse(layout, "2019-09-06 9:00:00.000000000 +0000 UTC")

	s45, _ := calendar.Event{s2019_9_6_4, s2019_9_6_5}.ToEventNode()
	s48, _ := calendar.Event{s2019_9_6_4, s2019_9_6_8}.ToEventNode()
	s57, _ := calendar.Event{s2019_9_6_5, s2019_9_6_7}.ToEventNode()
	s68, _ := calendar.Event{s2019_9_6_6, s2019_9_6_8}.ToEventNode()
	s69, _ := calendar.Event{s2019_9_6_6, s2019_9_6_9}.ToEventNode()
	s78, _ := calendar.Event{s2019_9_6_7, s2019_9_6_8}.ToEventNode()

	comp := []struct {
		p   calendar.EventNode
		s   calendar.EventNode
		exp int
	}{
		{s45, s68, 1},  //s68 is greater
		{s68, s68, 0},  //s68 is equal to s68
		{s45, s48, -1}, //s58 is greater
		{s68, s45, -1}, //s45 is less
		{s68, s57, -1}, //s57 starts prior so it's less
		{s69, s78, 1},  //s78 is greater
	}

	//Run CompareTo operation for each pair and compare to expected value
	for _, c := range comp {
		val := c.p.CompareTo(&c.s)

		if val != c.exp {
			t.Errorf("Expected compareTo result of nodes %v and %v to be %v but was %v. Failure.", c.p, c.s, c.exp, val)
		}
	}
}

//Validate that an error is returned when the start or end times are in a non-chronological order
func TestToEventNode(t *testing.T) {

	//Define layout and time ranges for testing
	layout := "2006-01-02 15:04:05.000000000 +0000 UTC"

	//Time ranges Sept 6 2019 4 AM - 10 AM
	s2019_9_6_4, _ := time.Parse(layout, "2019-09-06 4:00:00.000000000 +0000 UTC")
	s2019_9_6_5, _ := time.Parse(layout, "2019-09-06 5:00:00.000000000 +0000 UTC")

	//Define passing and failure events
	s45 := calendar.Event{s2019_9_6_4, s2019_9_6_5}
	s54 := calendar.Event{s2019_9_6_5, s2019_9_6_4}

	//Define test array
	comp := []struct {
		p             calendar.Event
		errorExpected bool
	}{
		{s45, false},
		{s54, true},
	}

	//Test
	for _, test := range comp {
		_, err := test.p.ToEventNode()

		if test.errorExpected && err == nil {
			t.Errorf("Expected ToEventNode conversion of node <%v> to generate an error. Failure.", test.p)

		} else if !test.errorExpected && err != nil {
			t.Errorf("Expected ToEventNode conversion of node <%v> to not generate an error <%v>. Failure.", test.p, err)

		}
	}
}
