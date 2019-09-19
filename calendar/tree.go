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
			//log.Printf("Left doesn't exist. Adding %v to left of %v", c, node)
			node.Left = c
		} else {
			//log.Printf("Left already exists, moving into %v with %v\n", node.Left, c)
			err = Insert(node.Left, c)

			if err != nil {
				return err
			}
		}
	} else {
		//insert right
		if node.Right == nil {
			//log.Printf("Right doesn't exist. Adding %v to Right of %v", c, node)
			node.Right = c
		} else {
			//log.Printf("Right already exists, moving into %v with %v\n", node.Right, c)
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

	//Check if this candidate node overlaps the current node
	if candidate.Overlaps(*node) {
		overlap = append(overlap, Pair{Can: candidate.Ev, Node: node.Ev})
	}

	//If the current node has a left child and the start time of candidate occurs prior
	//to the node's start time, then go to the left child and search for any overlaps
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

//GetOverlappingEvents - returns a series of Pair structs containing overlapping events. Possible
//error can occur if the event is misformatted.
//ie. event.start > event.end
func GetOverlappingEvents(events ...Event) ([]Pair, error) {

	var root *EventNode
	overlap := []Pair{}

	for _, event := range events {
		if root == nil {
			//Set the root node once. If events is empty then an empty slice will be returned.
			t, _ := event.ToEventNode()
			root = &t
			continue
		}

		//Covert event to an EventNode in order to allow left/right child relationships
		en, _ := event.ToEventNode()

		//Traverse tree and find any nodes that overlap the event that is going to added. Prevents duplicate [A,B] [B,A] pairs
		result := getOverlap(root, &en)

		//Append the overlapping pairs to the list
		overlap = append(overlap, result...)

		//Safe to insert the tree.
		if err := Insert(root, &en); err != nil {
			return nil, err
		}
	}

	return overlap, nil
}
