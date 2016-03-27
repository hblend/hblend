package location

import . "gopkg.in/check.v1"

func (w *World) Test_Navigate_Dir_To_LocalComponent(c *C) {
	l := New()

	n := l.Navigate("services/auth.js")

	c.Assert(n.Dir, Equals, "services/")
	c.Assert(n.Filename, Equals, "auth.js")
}

func (w *World) Test_Navigate_Dir_To_AbsoluteUrl(c *C) {
	l := New()

	n := l.Navigate("http://example.com/b/logo.png")

	c.Assert(n.Filename, Equals, "logo.png")
	c.Assert(n.Dir, Equals, "http://example.com/b/")
}

func (w *World) Test_Navigate_Dir_To_RelativeComponent(c *C) {
	l := New().Navigate("services/auth.js")

	n := l.Navigate("../controllers/cache.js")

	c.Assert(n.Filename, Equals, "cache.js")
	c.Assert(n.Dir, Equals, "controllers/")
}

func (w *World) Test_Navigate_Remote_To_RelativeComponent(c *C) {
	l := New().
		Navigate("https://example.com/lib/services/auth.js")

	n := l.Navigate("../controllers/cache.js")

	c.Assert(n.Filename, Equals, "cache.js")
	c.Assert(n.Dir, Equals, "https://example.com/lib/controllers/")
}

func (w *World) Test_IsRemote_Local(c *C) {
	l := New().
		Navigate("lib/services/auth.js")

	remote := l.IsRemote()

	c.Assert(remote, Equals, false)
}

func (w *World) Test_IsRemote_Remote(c *C) {
	l := New().
		Navigate("https://example.com/lib/services/auth.js")

	remote := l.IsRemote()

	c.Assert(remote, Equals, true)
}

func (w *World) Test_IsRemote_RelativeLocal(c *C) {
	l := New().
		Navigate("lib/services/auth.js").
		Navigate("../controllers/main.js")

	remote := l.IsRemote()

	c.Assert(remote, Equals, false)
}

func (w *World) Test_IsRemote_RelativeRemote(c *C) {
	l := New().
		Navigate("http://example.com/lib/services/auth.js").
		Navigate("../controllers/main.js")

	remote := l.IsRemote()

	c.Assert(remote, Equals, true)
}

func (w *World) Test_Canonical(c *C) {
	l := New().
		Navigate("lib/services/auth.js")

	canonical := l.Canonical()

	c.Assert(canonical, Equals, "lib/services/auth.js")
}
