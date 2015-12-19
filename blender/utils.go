package blender

import (
	"io/ioutil"
	"log"
)

func (this *Blender) ensure_imports(component string) {
	if exists, ok := this.included[component]; !ok || !exists {
		this.included[component] = true
		this.component_exists(component)

		css := this.blend_css(component)
		js := this.blend_js(component)

		this.CSS += css
		this.JS += js
	}
}

func (this *Blender) component_exists(component string) {
	_, err := ioutil.ReadDir(this.config.Dir.Components + "/" + component + "/")

	if nil != err {
		log.Println("Component `"+component+"` does not exist:", err)
	}
}
