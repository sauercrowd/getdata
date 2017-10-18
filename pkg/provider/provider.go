package provider

import (
	"time"
)

// Post is a generic struct for all platforms
type Post struct {
	ID         string    `json:"id"`
	User       string    `json:"user"`
	Content    string    `json:"content"`
	Votes      int64     `json:"votes"`
	Time       time.Time `json:"time"`
	LastUpdate time.Time `json:"lastupdate"`
}

// User is a generic struct for all platforms
type User struct {
	User    string `json:"user"`
	Monitor bool   `json:"monitor"`
}
