package entity

import "time"

type Task struct {
	ID        string        `json:"id"`
	Name      string        `json:"name"`
	Status    string        `json:"status"`
	CreatedAt time.Time     `json:"created"`
	StartedAt time.Time     `json:"started,omitempty"`
	Duration  time.Duration `json:"duration,omitempty"`
	InWork    time.Duration `json:"in_work,omitempty"`
	Result    string        `json:"result,omitempty"`
}
