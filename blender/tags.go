package blender

import (
	"encoding/base64"
	"hblend/utils"
	"html"
	"net/url"
	"path/filepath"
	"strings"

	"github.com/GerardoOscarJT/gotreescript"
)

func (this *Blender) tag_base64(component string, token *gotreescript.Token) string {

	path := token.Flags[0]
	filename := this.config.Dir.Components + "/" + component + "/" + path
	utils.CheckFileExists(filename)
	bytes := utils.ReadFileBytes(filename)
	return "data:;base64," + base64.StdEncoding.EncodeToString(bytes)
}

func (this *Blender) tag_path(prefix, component string, token *gotreescript.Token) string {

	path := token.Flags[0]
	filename := this.config.Dir.Components + "/" + component + "/" + path
	utils.CheckFileExists(filename)

	new_filename := utils.Md5File(filename) + filepath.Ext(filename)

	this.Files[filename] = PATH_FILES + new_filename

	return prefix + new_filename
}

func (this *Blender) tag_link(token *gotreescript.Token) string {

	component := token.Flags[0]

	if exists, ok := this.linked[component]; !ok || !exists {
		blender := NewBlender(this.config)
		blender.linked = this.linked
		blender.Blend(component)
	}

	return component + ".html"
}

func (this *Blender) tag_content(component string, token *gotreescript.Token) string {

	filename := this.config.Dir.Components + "/" + component + "/" + token.Flags[0]
	utils.CheckFileExists(filename)

	content := utils.ReadFile(filename)

	if !in_array_lowercase("no-parse", token.Flags) {
		parse_content := gotreescript.Parse(content)

		processed_content := ""
		for _, token := range *parse_content {
			if gotreescript.TEXT == token.Type {
				processed_content += token.Partial
			} else if gotreescript.TAG == token.Type {
				name := strings.ToLower(token.Name)
				if "include" == name {
					this.ensure_imports(token.Flags[0])
				} else if "path" == name {
					processed_content += this.tag_path(PATH_FILES, component, token)
				} else if "base64" == name {
					processed_content += this.tag_base64(component, token)
				} else if "content" == name {
					processed_content = this.tag_content(component, token)
				} else {
					processed_content += token.Partial
				}
			}
		}

		content = processed_content
	}

	escape, escape_ok := token.Args["escape"]
	if escape_ok {
		if "string" == escape {
			return strings.NewReplacer(
				"\\", "\\\\",
				"'", "\\'",
				"\"", "\\\"",
				"\n", "\\n",
				"\t", "\\t",
			).Replace(content)
		} else if "urlencode" == escape {
			return url.QueryEscape(content)
		} else if "html" == escape {
			return html.EscapeString(content)
		}
	}

	return content
}

func (this *Blender) tag_css_tags() string {

	msg := ""

	for _, filename := range this.CSS_tags {
		msg += "<link rel=\"stylesheet\" type=\"text/css\" href=\"" + filename + "\" title=\"default\">"
	}

	return msg
}

func (this *Blender) tag_js_tags() string {

	msg := ""

	for _, filename := range this.JS_tags {
		msg += "<script src=\"" + filename + "\" type=\"text/javascript\"></script>"
	}

	return msg
}
