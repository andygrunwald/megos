package mesos

import (
	"fmt"
)

func (c *Client) GetTaskByID(framework *Framework, taskID string) (*Task, error) {
	for _, task := range framework.CompletedTasks {
		if taskID == task.ID {
			return &task, nil
		}
	}

	return nil, fmt.Errorf("No task with id \"%s\" found in framework \"%s\"", taskID, framework.Name)
}
