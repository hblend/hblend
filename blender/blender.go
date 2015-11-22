package blender

import (
	"hblend/configuration"
	"hblend/utils"
	"io/ioutil"
	"log"
	"strings"

	"github.com/GerardoOscarJT/gotreescript"
)

type Blender struct {
	JS       string
	JS_tags  map[string]string
	CSS      string
	CSS_tags map[string]string
	HTML     string
	Files    map[string]string
	included map[string]bool
	config   *configuration.Configuration
}

func NewBlender(config *configuration.Configuration) *Blender {

	return &Blender{
		config:   config,
		included: map[string]bool{},
		Files:    map[string]string{},
		JS_tags:  map[string]string{},
		CSS_tags: map[string]string{},
	}
}

func (this *Blender) Blend(component string) {

	this.HTML = this.blend_html(component)
}

func (this *Blender) blend_html(component string) string {

	source_html := utils.ReadFile(this.config.Dir.Components + "/" + component + "/index.html")

	map_html := map[*gotreescript.Token]string{}

	parse_html := gotreescript.Parse(source_html)

	// Pass 1
	for _, token := range *parse_html {
		if gotreescript.TEXT == token.Type {
			map_html[token] = token.Partial
		} else if gotreescript.TAG == token.Type {
			name := strings.ToLower(token.Name)
			if "include" == name {
				this.ensure_imports(token.Flags[0])
				map_html[token] = this.blend_html(token.Flags[0])
			} else if "base64" == name {
				map_html[token] += this.tag_base64(component, token)
			} else if "file" == name {
				map_html[token] = this.tag_file(component, token)
			}
		}
	}

	this.ensure_imports(component)

	this.CSS_tags[""] = "files/" + utils.Md5String(this.CSS) + ".css"
	this.JS_tags[""] = "files/" + utils.Md5String(this.JS) + ".js"

	// Pass 2
	for _, token := range *parse_html {
		if gotreescript.TAG == token.Type {
			name := strings.ToLower(token.Name)
			if "js-tags" == name {
				map_html[token] = this.tag_js_tags()
			} else if "css-tags" == name {
				map_html[token] = this.tag_css_tags()
			}

		}
	}

	html := ""
	for _, token := range *parse_html {
		html += map_html[token]
	}

	return html
}

func (this *Blender) ensure_imports(component string) {
	if exists, ok := this.included[component]; !ok || !exists {
		this.included[component] = true
		this.component_exists(component)

		css := this.blend_css(component)
		js := this.blend_js(component)

		this.CSS += css
		this.JS += js
	}
}

func (this *Blender) component_exists(component string) {
	_, err := ioutil.ReadDir(this.config.Dir.Components + "/" + component + "/")

	if nil != err {
		log.Println("Component `"+component+"` does not exist:", err)
	}
}

func (this *Blender) blend_css(component string) string {

	source_css := utils.ReadFile(this.config.Dir.Components + "/" + component + "/index.css")

	parse_css := gotreescript.Parse(source_css)

	css := ""

	for _, token := range *parse_css {
		if gotreescript.TEXT == token.Type {
			css += token.Partial
		} else if gotreescript.TAG == token.Type {
			name := strings.ToLower(token.Name)
			if "include" == name {
				this.ensure_imports(token.Flags[0])
			} else if "file" == name {
				css += this.tag_file(component, token)
			} else if "base64" == name {
				css += this.tag_base64(component, token)
			}
		}
	}

	return css
}

func (this *Blender) blend_js(component string) string {

	source_js := utils.ReadFile(this.config.Dir.Components + "/" + component + "/index.js")

	parse_js := gotreescript.Parse(source_js)

	js := ""

	for _, token := range *parse_js {
		if gotreescript.TEXT == token.Type {
			js += token.Partial
		} else if gotreescript.TAG == token.Type {
			name := strings.ToLower(token.Name)
			if "include" == name {
				this.ensure_imports(token.Flags[0])
			} else if "file" == name {
				js += this.tag_file(component, token)
			}
		}
	}

	return js
}
