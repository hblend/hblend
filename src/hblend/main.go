package main

import (
	"hblend/blender"
	config "hblend/configuration"
	"hblend/utils"
)

func main() {

	b := blender.NewBlender()
	b.Blend(config.Component)

	for dst, src := range b.Files {
		if !utils.FileExists(dst) {
			utils.CopyFile(src, config.DirWww+"/"+dst)
		}
	}

	// _ = c
	// fmt.Println("HTML:\n" + *c.Html)
	// fmt.Println("CSS:\n" + *c.Css)
	// fmt.Println("JS:\n" + *c.Js)
}
