package blender

import (
	"fmt"
	"strings"

	"hblend/component"
	config "hblend/configuration"
	"hblend/location"
	"hblend/storage"
	"hblend/utils"
)

type Blender struct {
	Files    map[string]string
	linked   map[string]bool
	Storage  storage.Storager
	Location *location.Location
}

func New() *Blender {

	return &Blender{
		Files:    map[string]string{},
		linked:   map[string]bool{},
		Location: location.New(),
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

func (b *Blender) blend_component(name string) *component.Component {

	fmt.Printf("Blending %s\n", name)

	l := b.Location.Navigate(name)
	c := component.New(l)
	c.Storage = b.Storage
	c.Files = b.Files
	c.Blend()

	dst := config.DirWww + "/" + strings.Replace(c.Location.Canonical(), "/", "_", -1)
	utils.WriteFile(dst+".html", *c.Html)
	utils.WriteFile(dst+".css", *c.Css)
	utils.WriteFile(dst+".js", *c.Js)

	return c
}
