package maze

type Position struct {
	X int
	Y int
}

type Route []Position

func (s *Route) Push(p Position) {
	*s = append(*s, p)
}

func (s *Route) Pop() Position {
	// TODO len check
	l := len(*s)
	p := (*s)[l-1]
	*s = (*s)[:l-1]

	return p
}

// cell is an individual part of a maze with a route north, west, and a visited flag
// a "route" north, representing lack of a wall, is used instead of a wall=true as
// the default type for a bool is false so by default there is no route from that cell.
// This saves constructing cells with true values.
type Cell struct {
	NorthRoute bool
	WestRoute  bool
	Visited    bool
}

// maze represents the entire maze
type Maze struct {
	Cells [][]Cell
}

func New(width, height int, start Position) Maze {
	route := make(Route, 0)

	var m Maze
	m.Cells = make([][]Cell, width)

	for i := range m.Cells {
		m.Cells[i] = make([]Cell, height)
	}

	m.generate(start, &route)

	return m
}

func (m *Maze) generate(p Position, r *Route) {
	r.Push(p)
	c := &m.Cells[p.X][p.Y]
	c.Visited = true

	// TODO depth-first algorithm here

	// m.generate(next, &route)
}
