package entity

import "errors"

type Task struct {
	Id          string `json:"id"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

func NewTask(id string, description string, done bool) (*Task, error) {
	if done {
		return nil, errors.New("cannot create an already completed task")
	}
	return &Task{
		Id:          id,
		Description: description,
		Done:        done,
	}, nil
}
