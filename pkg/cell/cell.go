package cell

type Cell struct {
	X int
	Y int
}

func (c *Cell) GetPosition() (int, int) {
	return c.X, c.Y
}
