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
	s2019_9_6_4, _ := time.Parse(layout, "2010-09-06 4:00:00.000000000 +0000 UTC")
	s2019_9_6_5, _ := time.Parse(layout, "2019-09-06 5:00:00.000000000 +0000 UTC")
	s2019_9_6_6, _ := time.Parse(layout, "2019-09-06 6:00:00.000000000 +0000 UTC")
	s2019_9_6_8, _ := time.Parse(layout, "2019-09-06 8:00:00.000000000 +0000 UTC")

	e0, _ := calendar.Event{Start: s2019_9_6_4, End: s2019_9_6_8}.ToEventNode()
	e1, _ := calendar.Event{Start: s2019_9_6_6, End: s2019_9_6_8}.ToEventNode()
	e2, _ := calendar.Event{Start: s2019_9_6_4, End: s2019_9_6_5}.ToEventNode()

	//Build tree
	calendar.Insert(&e0, &e1)
	calendar.Insert(&e0, &e2)

	//Setup test stages. Expect to find e2 to the left and e1 to the right of e0
	setup := []struct {
		root           *calendar.EventNode
		node           *calendar.EventNode
		expectedOnLeft bool
	}{
		{&e0, &e2, true},
		{&e0, &e1, false},
	}

	for _, val := range setup {

		if err := calendar.Insert(val.root, val.node); err != nil {
			t.Errorf("Unexpected error when inserting node <%v> into root <%v>. %v", val.node, val.root, err)
		}

		//Check if expected to be on a particular side and that side is either null or not equal to expected node
		if (val.expectedOnLeft) && (val.root.Left == nil || (val.root.Left != val.node)) {
			t.Errorf("Expected node <%v> to be on left of root <%v>. Failure.", val.node, val.root)
		} else if (!val.expectedOnLeft) && (val.root.Right == nil || (val.root.Right != val.node)) {
			t.Errorf("Expected node <%v> to be on right of root <%v>. Failure.", val.node, val.root)
		}
	}

}
