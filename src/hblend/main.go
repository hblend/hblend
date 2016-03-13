package main

import (
	"flag"
	"fmt"

	"hblend/blender"
	. "hblend/constants"
	"hblend/utils"
)

func main() {

	flag.Parse()

	if 1 != len(flag.Args()) {
		fmt.Println("You should indicate the component you want, for example:\nhblend my-app")
		return
	}
	component := flag.Args()[0]

	b := blender.NewBlender()
	b.Blend(component)

	for dst, src := range b.Files {
		if !utils.FileExists(dst) {
			utils.CopyFile(src, DIR_WWW+"/"+dst)
		}
	}

	// _ = c
	// fmt.Println("HTML:\n" + *c.Html)
	// fmt.Println("CSS:\n" + *c.Css)
	// fmt.Println("JS:\n" + *c.Js)
}
