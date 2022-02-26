package layer

import "github.com/maxence-charriere/go-app/v9/pkg/app"

func Close(index int) {
	app.Window().Get("layer").Call("close", index)
}
