package layer

import "github.com/maxence-charriere/go-app/v9/pkg/app"

func Tip(text string, tag string) {
	app.Window().Get("layer").Call("tips", text, tag)
}
