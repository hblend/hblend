package component

import (
	"encoding/base64"
	"fmt"
	"html"
	"net/url"
	"path/filepath"
	"strings"

	"github.com/fulldump/gotreescript"

	config "hblend/configuration"
	"hblend/location"
	"hblend/storage"
	"hblend/utils"
)

type Component struct {
	Location *location.Location
	Storage  storage.Storager
	Html     *string
	Css      *string
	Js       *string
	included map[string]*Component
	Linked   map[string]bool
	Files    map[string]string
	TagsJs   []string
	TagsCss  []string
}

func New(l *location.Location) *Component {

	if "" == l.Filename {
		l.Filename = "index"
	}

	html := ""
	css := ""
	js := ""

	return &Component{
		Location: l,
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

	filename := "./" + c.Location.Filename + ".html"
	bytes, _ := c.Storage.ReadFileString(c.Location.Navigate(filename))
	parse_html := gotreescript.Parse(bytes)

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
	js := c.blend_js()
	css := c.blend_css()

	*c.Js += js
	*c.Css += css

	dst := c.Location.Navigate("./" + c.Location.Filename).Canonical()
	dst = strings.Replace(dst, "/", "_", -1)

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
				*c.Html += c.tag_path(token)
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

func (c *Component) blend_css() string {

	filename := "./" + c.Location.Filename + ".css"
	bytes, _ := c.Storage.ReadFileString(c.Location.Navigate(filename))
	parse_css := gotreescript.Parse(bytes)

	css := ""
	for _, token := range *parse_css {
		if gotreescript.TEXT == token.Type {
			css += token.Partial
		} else if gotreescript.TAG == token.Type {
			n := strings.ToLower(token.Name)

			if "include" == n {
				c.tag_include(token)
			} else if "base64" == n {
				css += c.tag_base64(token)
			} else if "content" == n {
				css += c.tag_content(token)
			} else if "path" == n {
				css += c.tag_path(token)
			} else if "link" == n {
				css += c.tag_link(token)
			} else if "todo:" == n {
				css += c.tag_todo(token)
			} else {
				css += c.tag_else(token)
			}
		}
	}

	return css
}

func (c *Component) blend_js() string {

	filename := "./" + c.Location.Filename + ".js"
	bytes, _ := c.Storage.ReadFileString(c.Location.Navigate(filename))
	parse_js := gotreescript.Parse(bytes)

	js := ""
	for _, token := range *parse_js {
		if gotreescript.TEXT == token.Type {
			js += token.Partial
		} else if gotreescript.TAG == token.Type {
			n := strings.ToLower(token.Name)

			if "include" == n {
				c.tag_include(token)
			} else if "base64" == n {
				js += c.tag_base64(token)
			} else if "content" == n {
				js += c.tag_content(token)
			} else if "path" == n {
				js += c.tag_path(token)
			} else if "link" == n {
				js += c.tag_link(token)
			} else if "todo:" == n {
				js += c.tag_todo(token)
			} else {
				js += c.tag_else(token)
			}
		}
	}

	return js
}

func (c *Component) Blend() {

	// TODO: Check if component does not exist

	c.blend_html()
}

/**
 * include once necessary css and js AND return the component
 */
func (c *Component) require(token *gotreescript.Token) *Component {

	component := token_component(token)

	l := c.Location.Navigate(component)

	canonical := l.Canonical()

	if item, exists := c.included[canonical]; exists {
		return item
	}

	n := New(l)
	n.Location = l
	n.Storage = c.Storage
	n.Css = c.Css
	n.Js = c.Js
	n.included = c.included
	n.Linked = c.Linked
	n.Files = c.Files

	n.included[canonical] = n

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
	bytes, _ := c.Storage.ReadFileBytes(c.Location.Navigate(filename))

	return "data:;base64," + base64.StdEncoding.EncodeToString(bytes)
}

func (c *Component) tag_content(token *gotreescript.Token) string {

	filename := token_component(token)

	content, _ := c.Storage.ReadFileString(c.Location.Navigate(filename))

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
					processed_content += c.tag_path(token)
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

func (c *Component) tag_path(token *gotreescript.Token) string {

	filename := token_component(token)

	l := c.Location.Navigate(filename)

	bytes, _ := c.Storage.ReadFileString(l)

	md5 := utils.Md5String(bytes)

	src := c.Storage.Path(l)
	dst := md5 + filepath.Ext(l.Filename)

	new_filename := config.DirFiles + "/" + dst

	c.Files[new_filename] = src

	return new_filename
}

func (c *Component) tag_link(token *gotreescript.Token) string {

	filename := token_component(token)

	link := c.Location.Navigate(filename)
	if "" == link.Filename {
		link.Filename = "index"
	}

	canonical := link.Canonical()
	c.Linked[canonical] = true

	canonical = strings.Replace(canonical, "/", "_", -1)

	return canonical + ".html"
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
