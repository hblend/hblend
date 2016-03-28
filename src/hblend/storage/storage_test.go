package storage

import (
	"hblend/location"

	. "gopkg.in/check.v1"
)

func (w *World) Test_ReadFileBytes_Local(c *C) {

	// Prepare
	w.WriteFileString("services/auth.js", "file auth.js")

	// Run
	l := location.New().Navigate("services/auth.js")
	bytes, err := w.Storage.ReadFileBytes(l)

	// Check
	c.Assert(string(bytes), Equals, "file auth.js")
	c.Assert(err, IsNil)
}

func (w *World) Test_ReadFileBytes_LocalError(c *C) {

	// Run
	l := location.New().Navigate("services/auth.js")
	bytes, err := w.Storage.ReadFileBytes(l)

	// Check
	c.Assert(string(bytes), Equals, "")
	c.Assert(err, NotNil)
}

func (w *World) Test_ReadFileBytes_RemoteError(c *C) {

	// Run
	l := location.New().Navigate("http://www.example.com/lib/services/ajax.js")
	bytes, err := w.Storage.ReadFileBytes(l)

	// Check
	c.Assert(string(bytes), Equals, "")
	c.Assert(err, NotNil)
}

func (w *World) Test_ReadFileBytes_RemoteCached(c *C) {

	// Prepare
	w.WriteFileString("remote/www.example.com/lib/services/ajax.js", "file ajax.js")

	// Run
	l := location.New().Navigate("http://www.example.com/lib/services/ajax.js")
	bytes, err := w.Storage.ReadFileBytes(l)

	// Check
	c.Assert(string(bytes), Equals, "file ajax.js")
	c.Assert(err, IsNil)
}

func (w *World) Test_ReadFileBytes_RemoteNoCached(c *C) {

	// Prepare
	w.WriteFileString("lib/services/remote.js", "file remote.js")

	// Run
	l := location.New().Navigate(w.Server.URL + "/lib/services/remote.js")
	bytes, err := w.Storage.ReadFileBytes(l)

	// Check
	c.Assert(string(bytes), Equals, "file remote.js")
	c.Assert(err, IsNil)
}

func (w *World) Test_ReadFileString_Ok(c *C) {

	// Prepare
	w.WriteFileString("services/common.js", "file common.js")

	// Run
	l := location.New().Navigate("services/common.js")
	bytes, err := w.Storage.ReadFileString(l)

	// Check
	c.Assert(bytes, Equals, "file common.js")
	c.Assert(err, IsNil)
}

func (w *World) Test_ReadFileString_Error(c *C) {

	// Run
	l := location.New().Navigate("services/common.js")
	bytes, err := w.Storage.ReadFileString(l)

	// Check
	c.Assert(bytes, Equals, "")
	c.Assert(err, NotNil)
}
