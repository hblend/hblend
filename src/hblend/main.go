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
		fmt.Println("You should indicate what component, for example:\nhblend my-app")
		return
	}
	component := flag.Args()[0]

	b := blender.NewBlender()
	c := b.Blend(component)

	for dst, src := range b.Files {
		utils.CopyFile(src, DIR_WWW+"/"+dst)
	}

	fmt.Println("HTML:\n" + *c.Html)
	fmt.Println("CSS:\n" + *c.Css)
	fmt.Println("JS:\n" + *c.Js)
}
