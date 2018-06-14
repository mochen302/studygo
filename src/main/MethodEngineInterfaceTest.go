package main

type Engine interface {
	Start()
	Stop()
}

type car struct {
	Engine
}

func (c *car) GoToWorkIn() {
	c.Start()
	c.Stop()
}
