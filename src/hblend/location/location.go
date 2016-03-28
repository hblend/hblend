package location

import (
	"net/url"
	"path/filepath"
	"strings"
)

type Location struct {
	Dir      string
	Filename string
}

func New() *Location {
	return &Location{}
}

func (l *Location) Navigate(path string) *Location {
	r := New()

	if is_remote(path) {
		u, _ := url.Parse(path)
		dir, filename := filepath.Split(u.Path)
		u.Path = dir

		r.Dir = u.String()
		r.Filename = filename
	} else if is_relative(path) {
		dir, filename := filepath.Split(path)

		if is_remote(l.Dir) {
			u, _ := url.Parse(l.Dir)
			u.Path = filepath.Join(u.Path, dir) + "/"
			r.Dir = u.String()
		} else {
			r.Dir = filepath.Join(l.Dir, dir) + "/"
		}

		r.Filename = filename
	} else {
		dir, filename := filepath.Split(path)

		r.Dir = dir
		r.Filename = filename
	}

	return r
}

func (l *Location) IsRemote() bool {
	return is_remote(l.Dir)
}

func (l *Location) Canonical() string {
	return filepath.Join(l.Dir, l.Filename)
}

func is_remote(path string) bool {
	path = strings.ToLower(path)

	is_http := strings.HasPrefix(path, "http://")
	is_https := strings.HasPrefix(path, "https://")

	return is_http || is_https
}

func is_relative(path string) bool {

	return strings.HasPrefix(path, ".")
}
