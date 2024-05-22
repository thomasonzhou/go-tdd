package sync

type Counter struct {
	Value int
}

func (c *Counter) Inc() {
	c.Value++
}
