package blender

// import (
// 	"strings"

// 	"github.com/fulldump/gotreescript"

// 	"hblend/utils"
// )

// func (this *Blender) blend_html(component string) string {

// 	source_html := utils.ReadFile(this.config.Dir.Components + "/" + component + "/index.html")

// 	map_html := map[*gotreescript.Token]string{}

// 	parse_html := gotreescript.Parse(source_html)

// 	// Pass 1
// 	for _, token := range *parse_html {
// 		if gotreescript.TEXT == token.Type {
// 			map_html[token] = token.Partial
// 		} else if gotreescript.TAG == token.Type {
// 			name := strings.ToLower(token.Name)
// 			if "include" == name {
// 				this.ensure_imports(token.Flags[0])
// 				map_html[token] = this.blend_html(token.Flags[0])
// 			} else if "base64" == name {
// 				map_html[token] += this.tag_base64(component, token)
// 			} else if "path" == name {
// 				map_html[token] = this.tag_path(PATH_FILES, component, token)
// 			} else if "content" == name {
// 				map_html[token] = this.tag_content(component, token)
// 			} else if "link" == name {
// 				map_html[token] = this.tag_link(token)
// 			}
// 		}
// 	}

// 	this.ensure_imports(component)

// 	this.CSS_tags[""] = PATH_FILES + utils.Md5String(this.CSS) + ".css"
// 	this.JS_tags[""] = PATH_FILES + utils.Md5String(this.JS) + ".js"

// 	// Pass 2
// 	for _, token := range *parse_html {
// 		if gotreescript.TAG == token.Type {
// 			name := strings.ToLower(token.Name)
// 			if "js-tags" == name {
// 				map_html[token] = this.tag_js_tags()
// 			} else if "css-tags" == name {
// 				map_html[token] = this.tag_css_tags()
// 			}

// 		}
// 	}

// 	html := ""
// 	for _, token := range *parse_html {
// 		html += map_html[token]
// 	}

// 	return html
// }

// func (this *Blender) blend_css(component string) string {

// 	source_css := utils.ReadFile(this.config.Dir.Components + "/" + component + "/index.css")

// 	parse_css := gotreescript.Parse(source_css)

// 	css := ""

// 	for _, token := range *parse_css {
// 		if gotreescript.TEXT == token.Type {
// 			css += token.Partial
// 		} else if gotreescript.TAG == token.Type {
// 			name := strings.ToLower(token.Name)
// 			if "include" == name {
// 				this.ensure_imports(token.Flags[0])
// 			} else if "path" == name {
// 				css += this.tag_path("", component, token)
// 			} else if "base64" == name {
// 				css += this.tag_base64(component, token)
// 			} else if "content" == name {
// 				css += this.tag_content(component, token)
// 			}
// 		}
// 	}

// 	return css
// }

// func (this *Blender) blend_js(component string) string {

// 	source_js := utils.ReadFile(this.config.Dir.Components + "/" + component + "/index.js")

// 	parse_js := gotreescript.Parse(source_js)

// 	js := ""

// 	for _, token := range *parse_js {
// 		if gotreescript.TEXT == token.Type {
// 			js += token.Partial
// 		} else if gotreescript.TAG == token.Type {
// 			name := strings.ToLower(token.Name)
// 			if "include" == name {
// 				this.ensure_imports(token.Flags[0])
// 			} else if "path" == name {
// 				js += this.tag_path(PATH_FILES, component, token)
// 			} else if "content" == name {
// 				js += this.tag_content(component, token)
// 			}
// 		}
// 	}

// 	return js
// }
