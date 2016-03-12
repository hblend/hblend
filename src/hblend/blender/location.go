package blender

import "strings"

type Location struct {
	Name   string // Canonical name
	Remote bool
	Schema string
}

func NewLocation(name string) *Location {
	l := &Location{
		Name: name,
	}

	if strings.HasPrefix(l.Name, "http://") {
		l.Remote = true
		l.Schema = "http"
		l.Name = strings.TrimPrefix(l.Name, "http://")

	} else if strings.HasPrefix(l.Name, "https://") {
		l.Remote = true
		l.Schema = "https"
		l.Name = strings.TrimPrefix(l.Name, "https://")

	} else {
		l.Remote = false
		l.Schema = "file"

	}

	l.Name = strings.Trim(l.Name, "/")

	return l
}
