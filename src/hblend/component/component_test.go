package component

import . "gopkg.in/check.v1"

func (w *World) Test_HelloWorld(c *C) {

	// Prepare
	w.WriteFile(w.LocalDir, "components/hello.html",
		`<h1>Hello World</h1>`,
	)
	w.WriteFile(w.LocalDir, "components/hello.css",
		`/* hello.css */`,
	)
	w.WriteFile(w.LocalDir, "components/hello.js",
		`/* hello.js */`,
	)

	// Run
	component := w.CreateComponent("components/hello")
	component.Blend()

	// Check
	c.Assert(*component.Html, Equals, `<h1>Hello World</h1>`)
	c.Assert(*component.Css, Equals, `/* hello.css */`)
	c.Assert(*component.Js, Equals, `/* hello.js */`)
}

func (w *World) Test_IncludeHtml(c *C) {

	// Prepare
	w.WriteFile(w.LocalDir, "components/hello.html",
		`<h1>[[include components/a]]</h1>`,
	)
	w.WriteFile(w.LocalDir, "components/hello.css",
		`/* hello.css */`,
	)
	w.WriteFile(w.LocalDir, "components/hello.js",
		`/* hello.js */`,
	)

	w.WriteFile(w.LocalDir, "components/a.html",
		`a.html`,
	)
	w.WriteFile(w.LocalDir, "components/a.css",
		`/* a.css */`,
	)
	w.WriteFile(w.LocalDir, "components/a.js",
		`/* a.js */`,
	)

	// Run
	component := w.CreateComponent("components/hello")
	component.Blend()

	// Check
	c.Assert(*component.Html, Equals, `<h1>a.html</h1>`)
	c.Assert(*component.Css, Equals, `/* a.css *//* hello.css */`)
	c.Assert(*component.Js, Equals, `/* a.js *//* hello.js */`)
}

func (w *World) Test_IncludeCss(c *C) {

	// Prepare
	w.WriteFile(w.LocalDir, "components/hello.html",
		`<h1>Hello World</h1>`,
	)
	w.WriteFile(w.LocalDir, "components/hello.css",
		`/* hello.css */[[include components/a]]/* hello.css (end) */`,
	)
	w.WriteFile(w.LocalDir, "components/hello.js",
		`/* hello.js */`,
	)

	w.WriteFile(w.LocalDir, "components/a.html",
		`a.html`,
	)
	w.WriteFile(w.LocalDir, "components/a.css",
		`/* a.css */`,
	)
	w.WriteFile(w.LocalDir, "components/a.js",
		`/* a.js */`,
	)

	// Run
	component := w.CreateComponent("components/hello")
	component.Blend()

	// Check
	c.Assert(*component.Html, Equals, `<h1>Hello World</h1>`)
	c.Assert(*component.Css, Equals, `/* a.css *//* hello.css *//* hello.css (end) */`)
	c.Assert(*component.Js, Equals, `/* a.js *//* hello.js */`)
}

func (w *World) Test_IncludeJs(c *C) {

	// Prepare
	w.WriteFile(w.LocalDir, "components/hello.html",
		`<h1>Hello World</h1>`,
	)
	w.WriteFile(w.LocalDir, "components/hello.css",
		`/* hello.css */`,
	)
	w.WriteFile(w.LocalDir, "components/hello.js",
		`/* hello.js */[[include components/a]]/* hello.js (end) */`,
	)

	w.WriteFile(w.LocalDir, "components/a.html",
		`a.html`,
	)
	w.WriteFile(w.LocalDir, "components/a.css",
		`/* a.css */`,
	)
	w.WriteFile(w.LocalDir, "components/a.js",
		`/* a.js */`,
	)

	// Run
	component := w.CreateComponent("components/hello")
	component.Blend()

	// Check
	c.Assert(*component.Html, Equals, `<h1>Hello World</h1>`)
	c.Assert(*component.Css, Equals, `/* a.css *//* hello.css */`)
	c.Assert(*component.Js, Equals, `/* a.js *//* hello.js *//* hello.js (end) */`)
}

func (w *World) Test_IncludeSlash_Fallback(c *C) {

	// Prepare
	w.WriteFile(w.LocalDir, "components/hello.html",
		`<h1>[[include components/a/]]</h1>`,
	)
	w.WriteFile(w.LocalDir, "components/hello.css",
		`/* hello.css */`,
	)
	w.WriteFile(w.LocalDir, "components/hello.js",
		`/* hello.js */`,
	)

	w.WriteFile(w.LocalDir, "components/a/index.html",
		`a.html`,
	)
	w.WriteFile(w.LocalDir, "components/a/index.css",
		`/* a.css */`,
	)
	w.WriteFile(w.LocalDir, "components/a/index.js",
		`/* a.js */`,
	)

	// Run
	component := w.CreateComponent("components/hello")
	component.Blend()

	// Check
	c.Assert(*component.Html, Equals, `<h1>a.html</h1>`)
	c.Assert(*component.Css, Equals, `/* a.css *//* hello.css */`)
	c.Assert(*component.Js, Equals, `/* a.js *//* hello.js */`)
}

func (w *World) Test_Base64_Local(c *C) {

	// Prepare
	w.WriteFile(w.LocalDir, "components/hello.html",
		`<img src="[[base64 ./image.png]]">`,
	)
	w.WriteFile(w.LocalDir, "components/hello.css",
		`b { background: url('[[base64 ./image.png]]'); }`,
	)
	w.WriteFile(w.LocalDir, "components/hello.js",
		`alert('[[base64 ./image.png]]')`,
	)

	w.WriteFile(w.LocalDir, "components/image.png",
		`--binary-image-content--`,
	)

	// Run
	component := w.CreateComponent("components/hello")
	component.Blend()

	// Check
	c.Assert(*component.Html, Equals, `<img src="data:;base64,LS1iaW5hcnktaW1hZ2UtY29udGVudC0t">`)
	c.Assert(*component.Css, Equals, `b { background: url('data:;base64,LS1iaW5hcnktaW1hZ2UtY29udGVudC0t'); }`)
	c.Assert(*component.Js, Equals, `alert('data:;base64,LS1iaW5hcnktaW1hZ2UtY29udGVudC0t')`)
}

func (w *World) Test_Base64_RemoteHtml(c *C) {

	// Prepare
	w.WriteFile(w.LocalDir, "components/hello.html",
		`<img src="[[base64 `+w.Server.URL+`/components/image.png]]">`,
	)
	w.WriteFile(w.LocalDir, "components/hello.css",
		`b { background: url('[[base64 `+w.Server.URL+`/components/image.png]]'); }`,
	)
	w.WriteFile(w.LocalDir, "components/hello.js",
		`alert('[[base64 `+w.Server.URL+`/components/image.png]]')`,
	)

	w.WriteFile(w.ServerDir, "components/image.png",
		`--binary-image-content--`,
	)

	// Run
	component := w.CreateComponent("components/hello")
	component.Blend()

	// Check
	c.Assert(*component.Html, Equals, `<img src="data:;base64,LS1iaW5hcnktaW1hZ2UtY29udGVudC0t">`)
	c.Assert(*component.Css, Equals, `b { background: url('data:;base64,LS1iaW5hcnktaW1hZ2UtY29udGVudC0t'); }`)
	c.Assert(*component.Js, Equals, `alert('data:;base64,LS1iaW5hcnktaW1hZ2UtY29udGVudC0t')`)

}

func (w *World) Test_Path_Local(c *C) {

	// Prepare
	w.WriteFile(w.LocalDir, "components/hello.html",
		`<img src="[[path ./image.png]]">`,
	)
	w.WriteFile(w.LocalDir, "components/hello.css",
		`b { background: url('[[path ./image.png]]'); }`,
	)
	w.WriteFile(w.LocalDir, "components/hello.js",
		`alert('[[path ./image.png]]')`,
	)

	w.WriteFile(w.LocalDir, "components/image.png",
		`--binary-image-content--`,
	)

	// Run
	component := w.CreateComponent("components/hello")
	component.Blend()

	// Check
	c.Assert(*component.Html, Equals, `<img src="files/75b95b20232c3d0675f597b3eadc6e00.png">`)
	c.Assert(*component.Css, Equals, `b { background: url('files/75b95b20232c3d0675f597b3eadc6e00.png'); }`)
	c.Assert(*component.Js, Equals, `alert('files/75b95b20232c3d0675f597b3eadc6e00.png')`)
}

func (w *World) Test_Path_Remote(c *C) {

	// Prepare
	w.WriteFile(w.LocalDir, "components/hello.html",
		`<img src="[[path `+w.Server.URL+`/components/image.png]]">`,
	)
	w.WriteFile(w.LocalDir, "components/hello.css",
		`b { background: url('[[path `+w.Server.URL+`/components/image.png]]'); }`,
	)
	w.WriteFile(w.LocalDir, "components/hello.js",
		`alert('[[path `+w.Server.URL+`/components/image.png]]')`,
	)

	w.WriteFile(w.ServerDir, "components/image.png",
		`--binary-image-content--`,
	)

	// Run
	component := w.CreateComponent("components/hello")
	component.Blend()

	// Check
	c.Assert(*component.Html, Equals, `<img src="files/75b95b20232c3d0675f597b3eadc6e00.png">`)
	c.Assert(*component.Css, Equals, `b { background: url('files/75b95b20232c3d0675f597b3eadc6e00.png'); }`)
	c.Assert(*component.Js, Equals, `alert('files/75b95b20232c3d0675f597b3eadc6e00.png')`)
}

func (w *World) Test_Todo(c *C) {

	// Prepare
	w.WriteFile(w.LocalDir, "components/hello.html",
		`<h1>Title</h1> [[TODO: complete index]]`,
	)
	w.WriteFile(w.LocalDir, "components/hello.css",
		`h1 {color: red;} [[TODO: add h2]]`,
	)
	w.WriteFile(w.LocalDir, "components/hello.js",
		`alert('Hello'); [[TODO: add user name]]`,
	)

	// Run
	component := w.CreateComponent("components/hello")
	component.Blend()

	// Check
	c.Assert(*component.Html, Equals, `<h1>Title</h1> `)
	c.Assert(*component.Css, Equals, `h1 {color: red;} `)
	c.Assert(*component.Js, Equals, `alert('Hello'); `)
}

func (w *World) Test_InventedTag(c *C) {

	// Prepare
	w.WriteFile(w.LocalDir, "components/hello.html",
		`[[inventedtaghtml]]`,
	)
	w.WriteFile(w.LocalDir, "components/hello.css",
		`[[inventedtagcss]]`,
	)
	w.WriteFile(w.LocalDir, "components/hello.js",
		`[[inventedtagjs]]`,
	)

	// Run
	component := w.CreateComponent("components/hello")
	component.Blend()

	// Check
	c.Assert(*component.Html, Equals, `[[inventedtaghtml]]`)
	c.Assert(*component.Css, Equals, `[[inventedtagcss]]`)
	c.Assert(*component.Js, Equals, `[[inventedtagjs]]`)
}

func (w *World) Test_JsTags(c *C) {

	// Prepare
	w.WriteFile(w.LocalDir, "components/hello.html",
		`[[js-tags]]`,
	)

	// Run
	component := w.CreateComponent("components/hello")
	component.Blend()

	// Check
	c.Assert(*component.Html, Equals, `<script src="components_hello.js" type="text/javascript"></script>`)
}

func (w *World) Test_CssTags(c *C) {

	// Prepare
	w.WriteFile(w.LocalDir, "components/hello.html",
		`[[css-tags]]`,
	)

	// Run
	component := w.CreateComponent("components/hello")
	component.Blend()

	// Check
	c.Assert(*component.Html, Equals, `<link rel="stylesheet" type="text/css" href="components_hello.css" title="default">`)
}

func (w *World) Test_Link(c *C) {

	// Prepare
	w.WriteFile(w.LocalDir, "components/hello.html",
		`<h1>Hello</h1> <a href="[[link ./bye]]">Bye</a>`,
	)
	w.WriteFile(w.LocalDir, "components/bye.html",
		`<h1>Bye</h1> <a href="[[link ./hello]]">Hello</a>`,
	)

	// Run
	component := w.CreateComponent("components/hello")
	component.Blend()

	// Check
	c.Assert(*component.Html, Equals, `<h1>Hello</h1> <a href="components_bye.html">Bye</a>`)
	c.Assert(component.Linked, DeepEquals, map[string]bool{
		"components/bye": true,
	})
}

func (w *World) Test_Link_Index(c *C) {

	// Prepare
	w.WriteFile(w.LocalDir, "components/hello/index.html",
		`<h1>Hello</h1> <a href="[[link components/bye/]]">Bye</a>`,
	)
	w.WriteFile(w.LocalDir, "components/bye/index.html",
		`<h1>Bye</h1> <a href="[[link components/hello/]]">Hello</a>`,
	)

	// Run
	component := w.CreateComponent("components/hello/")
	component.Blend()

	// Check
	c.Assert(*component.Html, Equals, `<h1>Hello</h1> <a href="components_bye_index.html">Bye</a>`)
	c.Assert(component.Linked, DeepEquals, map[string]bool{
		"components/bye/index": true,
	})
}

func (w *World) Test_Content_OK(c *C) {

	// Prepare
	w.WriteFile(w.LocalDir, "components/hello.js",
		`[[content lib/angular.js]]`,
	)
	w.WriteFile(w.LocalDir, "lib/angular.js",
		`/* This is the angular library */`,
	)

	// Run
	component := w.CreateComponent("components/hello")
	component.Blend()

	// Check
	c.Assert(*component.Html, Equals, ``)
	c.Assert(*component.Js, Equals, `/* This is the angular library */`)
}

func (w *World) Test_Content_NoParse(c *C) {

	// Prepare
	w.WriteFile(w.LocalDir, "components/hello.js",
		`[[content lib/angular.js no-parse]]`,
	)
	w.WriteFile(w.LocalDir, "lib/angular.js",
		`/* [[ This is the angular library ]] */`,
	)

	// Run
	component := w.CreateComponent("components/hello")
	component.Blend()

	// Check
	c.Assert(*component.Html, Equals, ``)
	c.Assert(*component.Js, Equals, `/* [[ This is the angular library ]] */`)
}

func (w *World) Test_Content_Escape(c *C) {

	// Prepare
	w.WriteFile(w.LocalDir, "components/hello.js",
		`[[content sample.txt]]
[[content sample.txt escape=html]]
[[content sample.txt escape=string]]
[[content sample.txt escape=urlencode]]`,
	)
	w.WriteFile(w.LocalDir, "sample.txt",
		`this <b>is</b> 'the local content'`,
	)

	// Run
	component := w.CreateComponent("components/hello")
	component.Blend()

	// Check
	c.Assert(*component.Html, Equals, ``)
	c.Assert(*component.Js, Equals, `this <b>is</b> 'the local content'
this &lt;b&gt;is&lt;/b&gt; &#39;the local content&#39;
this <b>is</b> \'the local content\'
this+%3Cb%3Eis%3C%2Fb%3E+%27the+local+content%27`)
}
