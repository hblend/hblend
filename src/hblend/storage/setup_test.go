package storage

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"path"
	"path/filepath"
	"testing"

	. "gopkg.in/check.v1"
)

type World struct {
	TempDir string
	Storage *Storage
	Server  *httptest.Server
}

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

func (w *World) SetUpSuite(c *C) {}

func (w *World) SetUpTest(c *C) {

	w.TempDir = c.MkDir()
	w.Server = httptest.NewServer(http.FileServer(http.Dir(w.TempDir)))

	config := &Config{
		BaseDir:   w.TempDir,
		RemoteDir: "remote",
	}

	w.Storage = New(config)
}

func (w *World) TearDownSuite(c *C) {}

func (w *World) WriteFileBytes(filename string, bytes []byte) {

	p := filepath.Join(w.TempDir, filename)

	// Ensure dir
	dir := path.Dir(p)
	if err := os.MkdirAll(dir, os.ModeDir|os.ModePerm); nil != err {
		panic(err)
	}
	// Ensure dir (end)

	ioutil.WriteFile(p, bytes, 0777)
}

func (w *World) WriteFileString(filename string, s string) {
	w.WriteFileBytes(filename, []byte(s))
}

var _ = Suite(&World{})
