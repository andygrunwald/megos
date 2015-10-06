package mesos

import (
	"fmt"
)

func (c *Client) GetSlaveByID(state *State, slaveID string) (*Slave, error) {
	for _, slave := range state.Slaves {
		if slaveID == slave.ID {
			return &slave, nil
		}
	}

	return nil, fmt.Errorf("No slave found with id \"%s\"", slaveID)
}
