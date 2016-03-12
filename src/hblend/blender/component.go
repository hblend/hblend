package blender

import (
	"encoding/base64"
	"fmt"
	"net/url"
	"path/filepath"
	"strings"

	"github.com/fulldump/gotreescript"

	. "hblend/constants"
	"hblend/utils"
	"html"
)

type Component struct {
	Location *Location
	Html     *string
	Css      *string
	Js       *string
	included map[string]*Component
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
		Files:    map[string]string{},
		TagsJs:   []string{},
		TagsCss:  []string{},
	}
}

func (c *Component) blend_html() {
	parse_html := gotreescript.Parse(c.ReadFileString("index.html"))

	// Traverse ONLY includes
	for _, token := range *parse_html {
		n := strings.ToLower(token.Name)
		if gotreescript.TAG == token.Type && "include" == n {
			c.require(token)
		}
	}

	c.blend_css()
	c.blend_js()

	c.TagsCss = append(c.TagsCss, DIR_FILES+"/"+utils.Md5String(*c.Css)+".css")
	c.TagsJs = append(c.TagsJs, DIR_FILES+"/"+utils.Md5String(*c.Js)+".js")

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
				*c.Html += c.tag_path(token, DIR_FILES+"/")
			} else if "css-tags" == n {
				*c.Html += c.tag_csstags(token)
			} else if "js-tags" == n {
				*c.Html += c.tag_jstags(token)
			} else {
				*c.Html += token.Partial
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
				*c.Css += c.tag_include(token)
			} else if "base64" == n {
				*c.Css += c.tag_base64(token)
			} else if "content" == n {
				*c.Css += c.tag_content(token)
			} else if "path" == n {
				*c.Css += c.tag_path(token, "")
			} else {
				*c.Css += token.Partial
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
				*c.Js += c.tag_include(token)
			} else if "base64" == n {
				*c.Js += c.tag_base64(token)
			} else if "content" == n {
				*c.Js += c.tag_content(token)
			} else if "path" == n {
				*c.Js += c.tag_path(token, DIR_FILES+"/")
			} else {
				*c.Js += token.Partial
			}
		}
	}
}

func (c *Component) Blend() {

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

	new_filename := DIR_FILES + "/" + dst

	c.Files[new_filename] = src
	// TODO: Add new_filename to a common files list

	return prefix + dst
}

/**
 * include necessary css and js
 */
func (c *Component) include_html(token *gotreescript.Token) {

	component := token_component(token)

	*c.Html += "[INCLUDE HTML FROM COMPONENT '" + component + "']"

}

func token_component(t *gotreescript.Token) string {
	if len(t.Flags) > 0 {
		return t.Flags[0]
	}

	if http, exists := t.Args["http"]; exists {
		return "http:" + http
	}

	if https, exists := t.Args["https"]; exists {
		return "https:" + https
	}

	return ""
}

func (c *Component) ReadPaths(filename string) string {
	l := NewLocation(filename)

	if l.Remote { // Absolute remote
		url := l.Schema + "://" + l.Name
		dst := DIR_COMPONENTS + "/" + l.Name
		if !utils.CheckFileExists(dst) {
			utils.CopyFileRemote(url, dst)
		}
		return dst
	}

	dst := DIR_COMPONENTS + "/" + c.Location.Name + "/" + l.Name

	if c.Location.Remote {
		url := c.Location.Schema + "://" + c.Location.Name + "/" + l.Name
		if !utils.CheckFileExists(dst) {
			if err := utils.CopyFileRemote(url, dst); nil != err {
				fmt.Println("WARNING: Fail downloading '%s': %s.\n", dst, err)
				return ""
			}
		}
	}

	return dst
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
