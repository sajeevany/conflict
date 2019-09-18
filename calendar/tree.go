package calendar

import "log"

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

//GetOverlap - gets the events that have overlapping times with the candidate
func GetOverlap(root, candidate *EventNode) []Pair {

	//var overlap []Pair
	overlap := []Pair{}

	if candidate.Overlaps(*root) {
		overlap = append(overlap, Pair{Can: *candidate, Node: *root})
	}

	if root.HasLeftChild() && candidate.Ev.Start.Before(root.Ev.Start) {
		overLappingPairs := GetOverlap(root.Left, candidate)
		overlap = append(overlap, overLappingPairs...)
	}

	if root.HasRightChild() && candidate.Ev.End.After(root.Ev.Start) {
		overLappingPairs := GetOverlap(root.Right, candidate)
		overlap = append(overlap, overLappingPairs...)
	}

	return overlap
}
