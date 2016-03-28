package configuration

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var DirComponents = "components"
var DirWww = "www"
var DirFiles = "files"
var DirRemote = "__remote__"
var Component = ""
var Verbose = false
var Help = false
var Clean = false

func Parse() {

	flag.StringVar(&DirWww, "output", "www", "Output directory")
	flag.StringVar(&DirRemote, "remote", "__remote__", "Remote directory")
	flag.BoolVar(&Clean, "clean", false, "Clean output directory")
	flag.BoolVar(&Verbose, "v", false, "Verbose: show extra information.")
	flag.BoolVar(&Help, "h", false, "Show this help")

	flag.Parse()

	// Show help and exit
	if Help {
		flag.PrintDefaults()
		os.Exit(0)
	}

	// Pick up component path
	if 1 != flag.NArg() {
		fmt.Println("You should indicate the component you want, for example:\nhblend project/index.html")
		os.Exit(-1)
	}
	DirComponents, Component = filepath.Split(flag.Arg(0))

	if "" == Component {
		Component = "index"
	}

	ext := filepath.Ext(Component)
	Component = strings.TrimSuffix(Component, ext)
	if "" == ext || "." == ext {
		ext = ".html"
	}

	ext_lower := strings.ToLower(ext)
	if ".html" != ext_lower && ".css" != ext_lower && ".js" != ext_lower {
		fmt.Println("You can not blend '" + ext + "' files, only .html, .css, .js")
		os.Exit(-2)
	}

	// If -v flag, print all configuration values
	if Verbose {
		fmt.Println("Configuration:")
		fmt.Printf("\tDirComponents: %s\n", DirComponents)
		fmt.Printf("\tDirWww: %s\n", DirWww)
		fmt.Printf("\tDirFiles: %s\n", DirFiles)
		fmt.Printf("\tComponent: %s\n", Component)
		fmt.Printf("\tVerbose: %t\n", Verbose)
	}

	if Clean {
		os.RemoveAll(DirWww)
	}

}
