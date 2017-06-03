package pathing

type Node struct {
	Tile                Vector2i
	Parent              *Node
	FCost, GCost, HCost float64
}

type ByFCost []Node

func NewNode(vec Vector2i, parent *Node, gCost float64, hCost float64) Node {
	return Node{
		Tile:   vec,
		Parent: parent,
		GCost:  gCost,
		HCost:  hCost,
		FCost:  gCost + hCost,
	}
}

func (node *Node) CompareTo(node2 *Node) int {
	if node.FCost > node2.FCost {
		return +1
	}
	if node.FCost < node2.FCost {
		return -1
	}
	return 0
}

//Comparators
func (a ByFCost) Len() int {
	return len(a)
}
func (a ByFCost) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
func (a ByFCost) Less(i, j int) bool {
	return a[i].FCost < a[j].FCost
}
