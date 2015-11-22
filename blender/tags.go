package blender

import (
	"encoding/base64"
	"hblend/utils"
	"path/filepath"

	"github.com/GerardoOscarJT/gotreescript"
)

func (this *Blender) tag_base64(component string, token *gotreescript.Token) string {

	path := token.Flags[0]
	filename := this.config.Dir.Components + "/" + component + "/" + path
	utils.CheckFileExists(filename)
	bytes := utils.ReadFileBytes(filename)
	return "data:;base64," + base64.StdEncoding.EncodeToString(bytes)
}

func (this *Blender) tag_file(component string, token *gotreescript.Token) string {

	path := token.Flags[0]
	filename := this.config.Dir.Components + "/" + component + "/" + path
	utils.CheckFileExists(filename)

	new_filename := "files/" + utils.Md5File(filename) + filepath.Ext(filename)

	this.Files[filename] = new_filename

	return new_filename
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
