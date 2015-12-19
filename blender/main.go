package blender

import (
	"hblend/configuration"
	"hblend/utils"
)

const PATH_FILES = "files/"

type Blender struct {
	JS       string
	JS_tags  map[string]string
	CSS      string
	CSS_tags map[string]string
	HTML     string
	Files    map[string]string
	included map[string]bool
	linked   map[string]bool
	config   *configuration.Configuration
}

func NewBlender(config *configuration.Configuration) *Blender {

	return &Blender{
		config:   config,
		included: map[string]bool{},
		Files:    map[string]string{},
		JS_tags:  map[string]string{},
		CSS_tags: map[string]string{},
		linked:   map[string]bool{},
	}
}

func (this *Blender) Blend(component string) {

	this.linked[component] = true

	this.HTML = this.blend_html(component)

	utils.WriteFile(this.config.Dir.Www+"/"+component+".html", this.HTML)
	utils.WriteFile(this.config.Dir.Www+"/"+this.CSS_tags[""], this.CSS)
	utils.WriteFile(this.config.Dir.Www+"/"+this.JS_tags[""], this.JS)

	for src, dst := range this.Files {
		utils.CopyFile(src, this.config.Dir.Www+"/"+dst)
	}
}
