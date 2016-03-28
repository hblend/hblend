package main

import (
	"hblend/blender"
	config "hblend/configuration"
	"hblend/storage"
	"hblend/utils"
)

func main() {

	config.Parse()

	b := blender.New()

	b.Storage = storage.New(&storage.Config{
		BaseDir:   config.DirComponents,
		RemoteDir: config.DirRemote,
	})

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
