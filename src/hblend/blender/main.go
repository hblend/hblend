package blender

import (
	"strings"

	. "hblend/constants"
	"hblend/utils"
)

type Blender struct {
	Files  map[string]string
	linked map[string]bool
}

func NewBlender() *Blender {

	return &Blender{
		Files:  map[string]string{},
		linked: map[string]bool{},
	}
}

func (b *Blender) Blend(name string) *Component {

	c := NewComponent(name)
	c.Files = b.Files
	c.Blend()

	utils.WriteFile(DIR_WWW+"/"+strings.Replace(c.Location.Name, "/", "_", -1)+".html", *c.Html)
	utils.WriteFile(DIR_WWW+"/"+DIR_FILES+"/"+utils.Md5String(*c.Css)+".css", *c.Css)
	utils.WriteFile(DIR_WWW+"/"+DIR_FILES+"/"+utils.Md5String(*c.Js)+".js", *c.Js)

	return c

	// this.linked[component] = true

}
