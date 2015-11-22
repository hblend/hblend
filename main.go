package main

import (
	"flag"
	"fmt"
	"hblend/blender"
	"hblend/configuration"
	"hblend/utils"
	"os"
)

var config = configuration.Configuration{}

func main() {

	flag.BoolVar(&config.ShowHelp, "help", false,
		"Show this help")
	flag.BoolVar(&config.Init, "init", false,
		"Initialize directories structure")

	flag.StringVar(&config.Dir.Components, "dir.components", "components",
		"Set components directory")
	flag.StringVar(&config.Dir.Www, "dir.www", "www",
		"Set www directory")

	flag.Parse()

	if config.ShowHelp {
		flag.Usage()
		return
	}

	if config.Init {
		initialize_directories()
		return
	}

	if 1 != len(flag.Args()) {
		fmt.Println("You should indicate what component, for example:\nhblend my-app")
		return
	}

	process(flag.Args()[0])
}

func initialize_directories() {
	os.Mkdir(config.Dir.Components, os.ModeDir|os.ModePerm)
}

func process(component string) {

	os.Mkdir(config.Dir.Www, os.ModeDir|os.ModePerm)
	os.Mkdir(config.Dir.Www+"/files", os.ModeDir|os.ModePerm)

	b := blender.NewBlender(&config)

	b.Blend(component)

	utils.WriteFile(config.Dir.Www+"/"+component+".html", b.HTML)
	utils.WriteFile(config.Dir.Www+"/"+b.CSS_tags[""], b.CSS)
	utils.WriteFile(config.Dir.Www+"/"+b.JS_tags[""], b.JS)

	for src, dst := range b.Files {
		utils.CopyFile(src, config.Dir.Www+"/"+dst)
	}
}
