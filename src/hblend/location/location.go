package location

import (
	"path/filepath"
	"strings"
)

type Location struct {
	Name     string // Canonical name
	Filename string
	Dir      string
	Remote   bool
	Schema   string
}

func NewLocation(name string) *Location {

	l := &Location{}

	if strings.HasPrefix(name, "http://") {
		l.Remote = true
		l.Schema = "http"
		name = strings.TrimPrefix(name, "http://")

	} else if strings.HasPrefix(name, "https://") {
		l.Remote = true
		l.Schema = "https"
		name = strings.TrimPrefix(name, "https://")
	} else {
		l.Remote = false
		l.Schema = "file"
		// name = name
	}

	l.Dir, l.Filename = filepath.Split(name)

	l.Name = l.Filename
	if "" == l.Name {
		l.Name = "index"
	}

	ext := filepath.Ext(l.Name)
	l.Name = strings.TrimSuffix(l.Name, ext)

	return l
}
