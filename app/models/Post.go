package models

type Post struct {
	ID
	Name        string
	Slug        string
	Description string
	TimeStamps
	SoftDeletes
}
