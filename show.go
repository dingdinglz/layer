package layer

import "github.com/maxence-charriere/go-app/v9/pkg/app"

func Msg(text string) {
	app.Window().Get("layer").Call("msg", text)
}
