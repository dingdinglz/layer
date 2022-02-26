package layer

import "github.com/maxence-charriere/go-app/v9/pkg/app"

func Load() int {
	return app.Window().Get("layer").Call("load").Int()
}

func Load1() int {
	return app.Window().Get("layer").Call("load", 1).Int()
}

func Load2() int {
	return app.Window().Get("layer").Call("load", 2).Int()
}
