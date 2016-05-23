package main

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/unixpickle/bezierfit"
)

func main() {
	js.Global.Set("onmessage", js.MakeFunc(handleMessage))
}

func handleMessage(this *js.Object, dataArg []*js.Object) interface{} {
	if len(dataArg) != 1 {
		panic("expected one argument")
	}

	data := dataArg[0].Get("data")
	requestID := data.Index(0)

	points := make([]bezierfit.Point, data.Length()-1)
	for i := 1; i < dataArg[0].Length(); i++ {
		arg := data.Index(i)
		x := arg.Get("x").Float()
		y := arg.Get("y").Float()
		points[i-1] = bezierfit.Point{X: x, Y: y}
	}

	anim := bezierfit.BestFit(points)
	resData := []interface{}{requestID, anim.P1.X, anim.P1.Y, anim.P2.X, anim.P2.Y}
	js.Global.Call("postMessage", resData)
	return nil
}
