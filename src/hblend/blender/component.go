package blender

import (
	"encoding/base64"
	"fmt"
	"net/url"
	"path/filepath"
	"strings"

	"github.com/fulldump/gotreescript"

	config "hblend/configuration"
	"hblend/utils"
	"html"
)

type Component struct {
	Location *Location
	Html     *string
	Css      *string
	Js       *string
	included map[string]*Component
	Linked   map[string]bool
	Files    map[string]string
	TagsJs   []string
	TagsCss  []string
}

func NewComponent(name string) *Component {
	html := ""
	css := ""
	js := ""

	return &Component{
		Location: NewLocation(name),
		Html:     &html,
		Css:      &css,
		Js:       &js,
		included: map[string]*Component{},
		Linked:   map[string]bool{},
		Files:    map[string]string{},
		TagsJs:   []string{},
		TagsCss:  []string{},
	}
}

func (c *Component) blend_html() {
	parse_html := gotreescript.Parse(c.ReadFileString("index.html"))

	// Traverse ONLY includes
	for _, token := range *parse_html {
		if gotreescript.TAG == token.Type {
			n := strings.ToLower(token.Name)
			if "include" == n {
				c.require(token)
				// } else if "content" == n {
				// 	c.tag_content(token)
			}
		}
	}

	// The order is very important! First js dependencies and Last css
	// dependencies (more priority)
	c.blend_js()
	c.blend_css()

	dst := strings.Replace(c.Location.Name, "/", "_", -1)

	c.TagsCss = append(c.TagsCss, dst+".css")
	c.TagsJs = append(c.TagsJs, dst+".js")

	for _, token := range *parse_html {
		if gotreescript.TEXT == token.Type {
			*c.Html += token.Partial
		} else if gotreescript.TAG == token.Type {
			n := strings.ToLower(token.Name)

			if "include" == n {
				*c.Html += c.tag_include(token)
			} else if "base64" == n {
				*c.Html += c.tag_base64(token)
			} else if "content" == n {
				*c.Html += c.tag_content(token)
			} else if "path" == n {
				*c.Html += c.tag_path(token, config.DirFiles+"/")
			} else if "link" == n {
				*c.Html += c.tag_link(token)
			} else if "css-tags" == n {
				*c.Html += c.tag_csstags(token)
			} else if "js-tags" == n {
				*c.Html += c.tag_jstags(token)
			} else if "todo:" == n {
				*c.Html += c.tag_todo(token)
			} else {
				*c.Html += c.tag_else(token)
			}
		}
	}

}

func (c *Component) blend_css() {
	parse_css := gotreescript.Parse(c.ReadFileString("index.css"))

	for _, token := range *parse_css {
		if gotreescript.TEXT == token.Type {
			*c.Css += token.Partial
		} else if gotreescript.TAG == token.Type {
			n := strings.ToLower(token.Name)

			if "include" == n {
				c.tag_include(token)
			} else if "base64" == n {
				*c.Css += c.tag_base64(token)
			} else if "content" == n {
				*c.Css += c.tag_content(token)
			} else if "path" == n {
				*c.Css += c.tag_path(token, config.DirFiles+"/")
			} else if "link" == n {
				*c.Css += c.tag_link(token)
			} else if "todo:" == n {
				*c.Css += c.tag_todo(token)
			} else {
				*c.Css += c.tag_else(token)
			}
		}
	}

}

func (c *Component) blend_js() {

	parse_js := gotreescript.Parse(c.ReadFileString("index.js"))

	for _, token := range *parse_js {
		if gotreescript.TEXT == token.Type {
			*c.Js += token.Partial
		} else if gotreescript.TAG == token.Type {
			n := strings.ToLower(token.Name)

			if "include" == n {
				c.tag_include(token)
			} else if "base64" == n {
				*c.Js += c.tag_base64(token)
			} else if "content" == n {
				*c.Js += c.tag_content(token)
			} else if "path" == n {
				*c.Js += c.tag_path(token, config.DirFiles+"/")
			} else if "link" == n {
				*c.Js += c.tag_link(token)
			} else if "todo:" == n {
				*c.Js += c.tag_todo(token)
			} else {
				*c.Js += c.tag_else(token)
			}
		}
	}

}

func (c *Component) Blend() {

	path := config.DirComponents + "/" + c.Location.Name
	if !utils.FileExists(path) {
		fmt.Printf("WARNING: Missing component '%s'.\n", c.Location.Name)
	}

	c.blend_html()
}

/**
 * include once necessary css and js AND return the component
 */
func (c *Component) require(token *gotreescript.Token) *Component {

	component := token_component(token)
	location := NewLocation(component)
	name := location.Name //  Normalized name

	if item, exists := c.included[name]; exists {
		return item
	}

	n := NewComponent(name)
	n.Location = location
	n.Css = c.Css
	n.Js = c.Js
	n.included = c.included
	n.Linked = c.Linked
	n.Files = c.Files

	n.included[name] = n

	n.Blend()

	return n
}

func (c *Component) tag_include(token *gotreescript.Token) string {
	s := c.require(token)

	return *s.Html
}

func (c *Component) tag_csstags(token *gotreescript.Token) string {
	s := ""
	for _, v := range c.TagsCss {
		s += "<link rel=\"stylesheet\" type=\"text/css\" href=\"" + v + "\" title=\"default\">"
	}
	return s
}

func (c *Component) tag_jstags(token *gotreescript.Token) string {
	s := ""
	for _, v := range c.TagsJs {
		s += "<script src=\"" + v + "\" type=\"text/javascript\"></script>"
	}
	return s
}

func (c *Component) tag_base64(token *gotreescript.Token) string {

	filename := token_component(token)
	bytes := c.ReadFile(filename)

	return "data:;base64," + base64.StdEncoding.EncodeToString(bytes)
}

func (c *Component) tag_content(token *gotreescript.Token) string {

	filename := token_component(token)
	content := c.ReadFileString(filename)

	if !utils.InArrayLowercase("no-parse", token.Flags) {
		parse_content := gotreescript.Parse(content)

		processed_content := ""
		for _, token := range *parse_content {
			if gotreescript.TEXT == token.Type {
				processed_content += token.Partial
			} else if gotreescript.TAG == token.Type {
				n := strings.ToLower(token.Name)
				if "include" == n {
					c.tag_include(token)
				} else if "base64" == n {
					processed_content += c.tag_base64(token)
				} else if "content" == n {
					processed_content = c.tag_content(token)
				} else if "path" == n {
					processed_content += c.tag_path(token, config.DirFiles+"/")
				} else if "link" == n {
					processed_content += c.tag_link(token)
				} else if "todo:" == n {
					processed_content += c.tag_todo(token)
				} else {
					processed_content += c.tag_else(token)
				}
			}
		}

		content = processed_content
	}

	escape, escape_ok := token.Args["escape"]
	if escape_ok {
		if "string" == escape {
			content = strings.NewReplacer(
				"\\", "\\\\",
				"'", "\\'",
				"\"", "\\\"",
				"\n", "\\n",
				"\t", "\\t",
			).Replace(content)
		} else if "urlencode" == escape {
			content = url.QueryEscape(content)
		} else if "html" == escape {
			content = html.EscapeString(content)
		}
	}

	return content
}

func (c *Component) tag_path(token *gotreescript.Token, prefix string) string {

	filename := token_component(token)
	content := c.ReadFileString(filename)
	md5 := utils.Md5String(content)

	src := c.ReadPaths(filename)
	dst := md5 + filepath.Ext(filename)

	new_filename := config.DirFiles + "/" + dst

	c.Files[new_filename] = src

	return prefix + dst
}

func (c *Component) tag_link(token *gotreescript.Token) string {

	filename := token_component(token)
	c.Linked[filename] = true

	location := NewLocation(filename)
	name := location.Name // Normalized name

	return strings.Replace(name, "/", "_", -1) + ".html"
}

func (c *Component) tag_else(token *gotreescript.Token) string {

	fmt.Printf("WARNING: Invalid token '%s'.\n", token.Partial)

	return token.Partial
}

func (c *Component) tag_todo(token *gotreescript.Token) string {

	fmt.Println("TODO:", token.Partial)

	return ""
}

func token_component(t *gotreescript.Token) string {
	if http, exists := t.Args["http"]; exists {
		return "http:" + http
	}

	if https, exists := t.Args["https"]; exists {
		return "https:" + https
	}

	if len(t.Flags) > 0 {
		return t.Flags[0]
	}

	return ""
}

func (c *Component) ReadPaths(filename string) string {
	l := NewLocation(filename)

	if l.Remote { // Absolute remote
		src := l.Schema + "://" + l.Name
		dst := config.DirComponents + "/" + l.Name
		download(src, dst)
		return dst
	}

	dst := config.DirComponents + "/" + c.Location.Name + "/" + l.Name

	if c.Location.Remote {
		src := c.Location.Schema + "://" + c.Location.Name + "/" + l.Name
		download(src, dst)
	}

	return dst
}

func download(src, dst string) {

	if !utils.FileExists(dst) {
		if err := utils.CopyFileRemote(src, dst); nil != err {
			fmt.Printf("WARNING: %s.\n", err)
			return
		}
		fmt.Printf("Downloading %s...OK\n", src)
		return
	}

	// fmt.Printf("Downloading %s...CACHE HIT\n", src)
}

/**
 * Read a file (as []byte) inside a component
 */
func (c *Component) ReadFile(filename string) []byte {

	dst := c.ReadPaths(filename)

	return utils.ReadFileBytes(dst)
}

/**
 * Read a file (as string) inside a component
 */
func (c *Component) ReadFileString(filename string) string {
	return string(c.ReadFile(filename))
}
