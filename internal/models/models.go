package models

import "time"

// Item ...
type Item struct {
	ID          int
	Name        string
	Description string
	Created     time.Time
	Updated     time.Time
	Deleted     time.Time
}
