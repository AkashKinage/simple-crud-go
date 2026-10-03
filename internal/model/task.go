package model

import "time"

var Statuses = []string{"pending", "in_progress", "done"}
var Priorities = []string{"low", "medium", "high"}

type Task struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	Priority    string    `json:"priority"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func IsValidStatus(s string) bool {
	for _, v := range Statuses {
        if s == v {
            return true
        }
    }
    return false
}

func IsValidPriority(s string) bool {
	for _, v := range Priorities {
        if s == v {
            return true
        }
    }
    return false
}