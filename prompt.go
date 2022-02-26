package layer

import "github.com/maxence-charriere/go-app/v9/pkg/app"

func Prompt(callback func(what string), title string) {
	app.Window().Get("layer").Call("prompt", map[string]interface{}{
		"title":      title,
		"shadeClose": true,
		"closeBtn":   0,
	}, app.FuncOf(func(this app.Value, args []app.Value) interface{} {
		callback(args[0].String())
		Close(args[1].Int())
		return nil
	}))
}
