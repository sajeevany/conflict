package calendar

import (
	"errors"
	"fmt"
	"log"
	"time"
)

//Event - Calendar event
type Event struct {
	Start time.Time
	End   time.Time
}

//EventNode - Event as node
type EventNode struct {
	Ev    Event
	Left  *EventNode
	Right *EventNode
}

//ToEventNode - Converts an event struct to an EventNode
func (e Event) ToEventNode() (EventNode, error) {

	if e.End.Before(e.Start) {
		msg := fmt.Sprintf("The specified end time <%v>  occurs before the start time <%v>.\n", e.End, e.Start)
		log.Println(msg)
		return EventNode{}, errors.New(msg)
	}

	return EventNode{Ev: e}, nil
}

//Overlaps - returns true if the time frame overlaps
func (e Event) Overlaps(c Event) bool {
	return !(e.Start.After(c.End) || e.End.Before(c.Start))
}

//Overlaps - returns true if the time frame overlaps
func (e EventNode) Overlaps(c EventNode) bool {
	return e.Ev.Overlaps(c.Ev)
}

//HasLeftChild - true if Left is non nil
func (e EventNode) HasLeftChild() bool {
	return e.Left != nil
}

//HasRightChild - true if Right is non nil
func (e EventNode) HasRightChild() bool {
	return e.Right != nil
}

/*CompareTo - returns a positive number if the eventNode being compared is
 *larger, 0 if equal, and a negative number if less.
 */
func (e EventNode) CompareTo(c *EventNode) int {

	/*
	 * Greater than -> c.low > e.low
	 * Less than -> c.high > e.high
	 * Equal -> cmp.equals
	 */

	//Candidate starts after
	if c.Ev.Start.After(e.Ev.Start) {
		return 1
	}

	//Candidate starts before
	if c.Ev.Start.Before(e.Ev.Start) {
		return -1
	}

	//Candidate starts at same time
	if c.Ev.Start.Equal(e.Ev.Start) {
		if c.Ev.End.Equal(e.Ev.End) {
			//Equal ending
			return 0
		}
		//ending is before/after
		return -1
	}

	//TODO Reevaluate this after writing unit tests. Compiler is issuing warning when not return past this point
	log.Fatalf("unexpected case with events %v and %v", e, c)
	return 0
}
