package calendar

import (
	"log"
)

//Pair - pair of EventNodes
type Pair struct {
	Can, Node Event
}

//Insert the candidate in the appropriate subtree of the node.
func Insert(node *EventNode, c *EventNode) error {

	cmp, err := node.CompareTo(c)

	if err != nil {
		log.Println(err)
		return err
	}

	if cmp < 1 {
		//insert left
		if node.Left == nil {
			log.Printf("Left doesn't exist. Adding %v to left of %v", c, node)
			node.Left = c
		} else {
			log.Printf("Left already exists, moving into %v with %v\n", node.Left, c)
			err = Insert(node.Left, c)

			if err != nil {
				return err
			}
		}
	} else {
		//insert right
		if node.Right == nil {
			log.Printf("Right doesn't exist. Adding %v to Right of %v", c, node)
			node.Right = c
		} else {
			log.Printf("Right already exists, moving into %v with %v\n", node.Right, c)
			err = Insert(node.Right, c)

			if err != nil {
				return err
			}
		}
	}

	return nil
}

//getOverlap - gets the events that have overlapping times with the candidate
func getOverlap(node, candidate *EventNode) []Pair {

	overlap := []Pair{}

	if candidate.Overlaps(*node) {
		overlap = append(overlap, Pair{Can: candidate.Ev, Node: node.Ev})
	}

	if node.HasLeftChild() && candidate.Ev.Start.Before(node.Ev.Start) {
		overLappingPairs := getOverlap(node.Left, candidate)
		overlap = append(overlap, overLappingPairs...)
	}

	if node.HasRightChild() && candidate.Ev.End.After(node.Ev.Start) {
		overLappingPairs := getOverlap(node.Right, candidate)
		overlap = append(overlap, overLappingPairs...)
	}

	return overlap
}

//GetOverlappingEvents - returns a series of Pair structs containing overlapping events
func GetOverlappingEvents(events ...Event) []Pair {

	var root *EventNode
	overlap := []Pair{}

	for _, event := range events {
		if root == nil {
			t, _ := event.ToEventNode()
			root = &t
			continue
		}

		en, _ := event.ToEventNode()
		result := getOverlap(root, &en)
		overlap = append(overlap, result...)
		Insert(root, &en)
	}

	return overlap
}
