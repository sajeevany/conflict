package calendar

import (
	"log"
)

//Pair - pair of EventNodes
type Pair struct {
	Can, Node EventNode
}

//Insert the candidate in the appropriate subtree of the node.
func Insert(node *EventNode, c *EventNode) {

	if node.CompareTo(c) < 1 {
		//insert left
		if node.Left == nil {
			log.Printf("Left doesn't exist. Adding %v to left of %v", c, node)
			node.Left = c
		} else {
			log.Printf("Left already exists, moving into %v with %v\n", node.Left, c)
			Insert(node.Left, c)
		}
	} else {
		//insert right
		if node.Right == nil {
			log.Printf("Right doesn't exist. Adding %v to Right of %v", c, node)
			node.Right = c
		} else {
			log.Printf("Right already exists, moving into %v with %v\n", node.Right, c)
			Insert(node.Right, c)
		}
	}
}

//getOverlap - gets the events that have overlapping times with the candidate
func getOverlap(root, candidate *EventNode) []Pair {

	//var overlap []Pair
	overlap := []Pair{}

	if candidate.Overlaps(*root) {
		overlap = append(overlap, Pair{Can: *candidate, Node: *root})
	}

	if root.HasLeftChild() && candidate.Ev.Start.Before(root.Ev.Start) {
		overLappingPairs := getOverlap(root.Left, candidate)
		overlap = append(overlap, overLappingPairs...)
	}

	if root.HasRightChild() && candidate.Ev.End.After(root.Ev.Start) {
		overLappingPairs := getOverlap(root.Right, candidate)
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
			//Look into this
			t := event.ToEventNode()
			root = &t
			continue
		}

		en := event.ToEventNode()
		result := getOverlap(root, &en)
		overlap = append(overlap, result...)
		Insert(root, &en)
	}

	return overlap
}
