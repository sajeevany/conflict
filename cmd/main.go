package main

import (
	"fmt"
	"log"
	"time"

	"github.com/sajeevany/conflict/calendar"
)

func main() {
	//Sample execution
	//Define layout and time ranges for testing
	layout := "2006-01-02 15:04:05.000000000 +0000 UTC"

	//Time ranges Sept 6 2019 4 AM - 10 AM
	s2019_9_6_4, _ := time.Parse(layout, "2019-09-06 4:00:00.000000000 +0000 UTC")
	s2019_9_6_5, _ := time.Parse(layout, "2019-09-06 5:00:00.000000000 +0000 UTC")
	s2019_9_6_6, _ := time.Parse(layout, "2019-09-06 6:00:00.000000000 +0000 UTC")
	s2019_9_6_8, _ := time.Parse(layout, "2019-09-06 8:00:00.000000000 +0000 UTC")
	s2019_9_6_7, _ := time.Parse(layout, "2019-09-06 7:00:00.000000000 +0000 UTC")
	s2019_9_6_10, _ := time.Parse(layout, "2019-09-06 10:00:00.000000000 +0000 UTC")

	//Define some events based on above times
	e0 := calendar.Event{Start: s2019_9_6_6, End: s2019_9_6_8}
	e1 := calendar.Event{Start: s2019_9_6_4, End: s2019_9_6_8}
	e2 := calendar.Event{Start: s2019_9_6_4, End: s2019_9_6_5}
	e3 := calendar.Event{Start: s2019_9_6_6, End: s2019_9_6_7}
	e4 := calendar.Event{Start: s2019_9_6_7, End: s2019_9_6_10}
	e5 := calendar.Event{Start: s2019_9_6_4, End: s2019_9_6_10}

	//Get the overlapping events as pairs
	overlaps, err := calendar.GetOverlappingEvents([]calendar.Event{e0, e1, e2, e3, e4, e5}...)

	//Check for errors
	if err != nil {
		log.Println(fmt.Sprintf("An error occurred while finding overlapping events. <%v>", err))
	} else {
		fmt.Println("Overlapping events: ")
		for _, i := range overlaps {
			fmt.Printf("Pair %v %v\n", i.Can, i.Node)
		}
	}
}
