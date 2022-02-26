package layer

import "github.com/maxence-charriere/go-app/v9/pkg/app"

func Alert_Success(title string, text string) {
	app.Window().Get("layer").Call("alert", text, map[string]interface{}{
		"icon":       1,
		"title":      title,
		"shadeClose": true,
		"closeBtn":   0,
	})
}

func Alert_Fail(title string, text string) {
	app.Window().Get("layer").Call("alert", text, map[string]interface{}{
		"icon":       2,
		"title":      title,
		"shadeClose": true,
		"closeBtn":   0,
	})
}

func Alert_Common(title string, text string) {
	app.Window().Get("layer").Call("alert", text, map[string]interface{}{
		"title":      title,
		"shadeClose": true,
		"closeBtn":   0,
	})
}
