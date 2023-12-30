package navigation

type Coordinate struct {
	Y int
	X int
}



type byXY []Coordinate

func (c byXY) Len() int {
	return len(c)
}
func (c byXY) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}
func (c byXY) Less(i, j int) bool {
	return c[i].Y < c[j].Y || (c[i].Y == c[j].Y && c[i].X < c[j].X)
}

func (c *Coordinate) Step(direction int, amount int) {

	for i := 0; i < amount; i++ {
		switch direction {
		case 0:
			c.Y--
		case 90:
			c.X++
		case 180:
			c.Y++
		case 270:
			c.X--
		}
	}

}