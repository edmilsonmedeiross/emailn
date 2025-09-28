package campaign

import "time"

type Contact struct {
	Email string
}

type Campaign struct {
	ID          int
	Name 			string
	CreatedOn time.Time
	Content   string
	Contacts []Contact
}