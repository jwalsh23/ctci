package data_structures

type IntGraph struct {
	vertexEdge map[int]*LinkedList
}

func (g *IntGraph) AddNode(val int) {
	if g.vertexEdge == nil {
		g.vertexEdge = map[int]*LinkedList{
			val: {},
		}
		return
	}
	if _, ok := g.vertexEdge[val]; !ok {
		g.vertexEdge[val] = &LinkedList{}
	}
}

func (g *IntGraph) AddRelation(v1, v2 int, bidirectional bool) {
	if _, ok := g.vertexEdge[v1]; !ok {
		g.AddNode(v1)
	}
	if _, ok := g.vertexEdge[v2]; !ok {
		g.AddNode(v2)
	}
	g.vertexEdge[v1].Add(v2)
	if bidirectional {
		g.vertexEdge[v2].Add(v1)
	}
}

func (g *IntGraph) RelationExist(v1, v2 int) bool {
	if _, ok := g.vertexEdge[v1]; !ok {
		return false
	}
	if _, ok := g.vertexEdge[v2]; !ok {
		return false
	}
	nodeQueue := Queue{}
	nodeQueue.Enqueue(v1)
	seen := map[int]struct{}{}
	return g.depthFirstSearch(v1, v2, seen)
	// return g.breadthFirstSearch(v2, &nodeQueue)
}

func (g *IntGraph) breadthFirstSearch(child int, nodeQueue *Queue) bool {
	parent := nodeQueue.Dequeue()
	if parent == 0 {
		return false
	}
	for _, node := range g.vertexEdge[parent].Slice() {
		if node == child {
			return true
		}
		nodeQueue.Enqueue(node)
	}
	val := g.breadthFirstSearch(child, nodeQueue)
	return val
}
func (g *IntGraph) depthFirstSearch(v1, v2 int, seen map[int]struct{}) bool {
	if len(g.vertexEdge[v1].Slice()) == 0 {
		return false
	} else {
		for _, val := range g.vertexEdge[v1].Slice() {
			if _, ok := seen[val]; ok {
				continue
			}
			seen[val] = struct{}{}
			found := val == v2 || g.depthFirstSearch(val, v2, seen)
			if found {
				return true
			}
		}
	}
	return false
}
