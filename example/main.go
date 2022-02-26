package main

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"net/http"
)

func main()  {
	app.Route("/",&Index{})
	app.RunWhenOnBrowser()
	http.Handle("/",&app.Handler{
		Name: "go-layer example",
		Description: "a example for dinglz's layer",
		Scripts: []string{
			"/web/layui/layui.js",
		},
		Styles: []string{
			"/web/layui/css/layui.css",
		},
	})
	http.ListenAndServe(":80",nil)
}
