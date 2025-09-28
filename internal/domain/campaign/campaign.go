package campaign

import "time"

type Campaign struct {
	ID          int
	Name 			string
	CreatedOn time.Time
	Content   string
	Contacts []string
}