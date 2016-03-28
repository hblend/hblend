package storage

import "hblend/location"

type Storager interface {
	ReadFileBytes(l *location.Location) ([]byte, error)
	ReadFileString(l *location.Location) (string, error)
	Path(l *location.Location) string
	Exists(l *location.Location) bool
}
