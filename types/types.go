package types

import "time"

type Reminder struct {
	Rem_Id      int           `json:"rem_id"`
	Name        string        `json:"name"`
	Duration    time.Duration `json:"duration"`
	IsRecursive bool          `json:"isRecursive"`
	IsActive    bool          `json:"isActive"`
}

type Query struct {
	IsActive    string
	IsRecursive string
}

type QueryDel struct {
	ID          int
	IsActive    string
	IsRecursive string
}
