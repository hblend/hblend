package storage

import (
	"fmt"
	"hblend/location"
	"hblend/utils"
	"io/ioutil"
	"net/url"
	"path/filepath"
)

type Storage struct {
	Config *Config
}

func New(c *Config) *Storage {
	return &Storage{
		Config: c,
	}
}

func (s *Storage) Path(l *location.Location) string {
	path := ""
	if l.IsRemote() {
		u, _ := url.Parse(l.Dir)
		path = filepath.Join(s.Config.BaseDir, s.Config.RemoteDir, u.Host, u.Path, l.Filename)
		download(l.Dir+l.Filename, path)
	} else {
		path = filepath.Join(s.Config.BaseDir, l.Dir, l.Filename)
	}

	return path
}

func (s *Storage) Exists(l *location.Location) bool {
	p := s.Path(l)
	return utils.FileExists(p)
}

func (s *Storage) ReadFileBytes(l *location.Location) ([]byte, error) {

	path := s.Path(l)

	bytes, bytes_err := ioutil.ReadFile(path)
	if nil != bytes_err {
		return []byte{}, bytes_err
	}

	return bytes, nil
}

func (s *Storage) ReadFileString(l *location.Location) (string, error) {
	bytes, err := s.ReadFileBytes(l)
	if nil != err {
		return "", err
	}

	return string(bytes), nil
}

func download(src, dst string) error {

	if !utils.FileExists(dst) {
		if err := utils.CopyFileRemote(src, dst); nil != err {
			fmt.Printf("WARNING: %s.\n", err)
			return err
		}
		fmt.Printf("Downloading %s...OK\n", src)
	}

	return nil
}
