package model

import "sync"

type connection struct {
	state string
}

func (c *connection) GetState() string {
	return c.state
}

func (c *connection) SetState(state string) {
	c.state = state
}

var connectionVar *connection
var once sync.Once

// function exportable
func CreateConnection() *connection {
	once.Do(func() {
		connectionVar = &connection{state: "off"}
	})
	return connectionVar
}
