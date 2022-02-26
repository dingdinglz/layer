package main

import (
	"gitee.com/dinglz/layer"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"time"
)

type Index struct {
	app.Compo
}

func (i *Index) Render() app.UI {
	return app.Div().Body(
		app.H1().Text("example of go-layer powered by 丁丁(dinglz)"),
		app.H2().Text("https://gitee.com/dinglz/layer"),
		app.Br(),
		app.Button().Text("msg").Class("layui-btn").OnClick(i.msgshow),
		app.Button().Text("alert").Class("layui-btn").OnClick(i.alertshow),
		app.Button().Text("alert-success").Class("layui-btn").OnClick(i.alertsshow),
		app.Button().Text("alert-fail").Class("layui-btn").OnClick(i.alertfshow),
		app.Button().Text("tip").Class("layui-btn").ID("tips").OnClick(i.tipshow),
		app.Br(),
		app.Hr(),
		app.H3().Text("loading"),
		app.Button().Text("load1").Class("layui-btn").OnClick(i.loading1),
		app.Button().Text("load2").Class("layui-btn").OnClick(i.loading2),
		app.Button().Text("load3").Class("layui-btn").OnClick(i.loading3),
		app.Br(),
		app.Hr(),
		app.H3().Text("Prompt"),
		app.Button().Text("Prompt").Class("layui-btn").OnClick(i.showp),
		app.Br(),
		app.Hr(),
		app.H3().Text("content"),
		app.Button().Text("联系作者").Class("layui-btn layui-btn-primary layui-border-blue").OnClick(i.cont),
	).Style("text-align","center")
}

func (i *Index)OnPreRender(ctx app.Context)  {
	ctx.Page().SetTitle("go-layer-example")
	ctx.Page().SetDescription("a example for go-layer")
}

func (i *Index) msgshow(ctx app.Context, e app.Event) {
	layer.Msg("hello,i am msg")
}

func (i *Index) loading1(ctx app.Context, e app.Event) {
	index:=layer.Load()
	go func() {
		time.Sleep(2*time.Second)
		layer.Close(index)
	}()
}

func (i *Index) loading2(ctx app.Context, e app.Event) {
	index:=layer.Load1()
	go func() {
		time.Sleep(2*time.Second)
		layer.Close(index)
	}()
}

func (i *Index) loading3(ctx app.Context, e app.Event) {
	index:=layer.Load2()
	go func() {
		time.Sleep(2*time.Second)
		layer.Close(index)
	}()
}

func (i *Index) cont(ctx app.Context, e app.Event) {
	app.Window().Get("location").Set("href","https://jq.qq.com/?_wv=1027&k=TdLRaYbp")
}

func (i *Index) alertshow(ctx app.Context, e app.Event) {
	layer.Alert_Common("普通","信息内容")
}

func (i *Index) alertsshow(ctx app.Context, e app.Event) {
	layer.Alert_Success("成功","信息内容")
}

func (i *Index) alertfshow(ctx app.Context, e app.Event) {
	layer.Alert_Fail("失败","信息内容")
}

func (i *Index) tipshow(ctx app.Context, e app.Event) {
	layer.Tip("我是tip","#tips")
}

func (i *Index) showp(ctx app.Context, e app.Event) {
	layer.Prompt(func(what string) {
		layer.Alert_Common("您输入的是",what)
	},"请输入一点东西")
}