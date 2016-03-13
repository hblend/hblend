package blender

import (
	"fmt"
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

func (b *Blender) Blend(name string) {

	b.add_component(name)
	b.blend_all()
}

func (b *Blender) add_component(name string) {

	if _, exist := b.linked[name]; !exist {
		b.linked[name] = false
		b.Blend(name)
	}
}

func (b *Blender) blend_all() {

	for name, linked := range b.linked {
		if !linked {
			b.linked[name] = true
			c := b.blend_component(name)
			for n, _ := range c.Linked {
				b.add_component(n)
			}
		}
	}
}

func (b *Blender) blend_component(name string) *Component {

	fmt.Printf("Blending %s\n", name)

	c := NewComponent(name)
	c.Files = b.Files
	c.Blend()

	utils.WriteFile(DIR_WWW+"/"+strings.Replace(c.Location.Name, "/", "_", -1)+".html", *c.Html)
	utils.WriteFile(DIR_WWW+"/"+DIR_FILES+"/"+utils.Md5String(*c.Css)+".css", *c.Css)
	utils.WriteFile(DIR_WWW+"/"+DIR_FILES+"/"+utils.Md5String(*c.Js)+".js", *c.Js)

	return c
}
