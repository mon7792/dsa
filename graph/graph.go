package graph

// gnode is representation of graph node.
type gnode struct {
	value string
	node  *gnode
}

// graphStore contains entire graph in the link list representation.
type graphStore struct {
	node *gnode
}

func NewGraph() *graphStore {
	return &graphStore{}
}
