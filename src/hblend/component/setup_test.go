package component

import (
	"hblend/location"
	"hblend/storage"
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
	LocalDir  string
	ServerDir string
	Server    *httptest.Server
	Storage   storage.Storager
}

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

func (w *World) SetUpSuite(c *C) {}

func (w *World) SetUpTest(c *C) {

	w.LocalDir = c.MkDir()
	w.ServerDir = c.MkDir()

	w.Server = httptest.NewServer(http.FileServer(http.Dir(w.ServerDir)))

	w.Storage = storage.New(&storage.Config{
		BaseDir:   w.LocalDir,
		RemoteDir: "remote",
	})
}

func (w *World) TearDownSuite(c *C) {}

func (w *World) CreateComponent(name string) *Component {
	l := location.New().Navigate(name)

	c := New(l)
	c.Storage = w.Storage

	return c
}

func (w *World) WriteFile(dir string, filename string, bytes string) {

	p := filepath.Join(dir, filename)

	// Ensure dir
	d := path.Dir(p)
	if err := os.MkdirAll(d, os.ModeDir|os.ModePerm); nil != err {
		panic(err)
	}
	// Ensure dir (end)

	ioutil.WriteFile(p, []byte(bytes), 0777)
}

var _ = Suite(&World{})
