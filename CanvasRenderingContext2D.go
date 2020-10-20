// +build js,wasm

/*
DO NOT EDIT

Code automatically generated by github.com/PapaCharlie/go-jsapi
*/
package jsapi

import js "syscall/js"

type CanvasRenderingContext2D js.Value

func (wrapper *CanvasRenderingContext2D) BeginPath() {
	(*js.Value)(wrapper).Call("beginPath")
}

func (wrapper *CanvasRenderingContext2D) Stroke() {
	(*js.Value)(wrapper).Call("stroke")
}

func (wrapper *CanvasRenderingContext2D) Fill() {
	(*js.Value)(wrapper).Call("fill")
}

func (wrapper *CanvasRenderingContext2D) FillRect(x float64, y float64, width float64, height float64) {
	(*js.Value)(wrapper).Call("fillRect", x, y, width, height)
}

func (wrapper *CanvasRenderingContext2D) StrokeRect(x float64, y float64, width float64, height float64) {
	(*js.Value)(wrapper).Call("strokeRect", x, y, width, height)
}

func (wrapper *CanvasRenderingContext2D) ClearRect(x float64, y float64, width float64, height float64) {
	(*js.Value)(wrapper).Call("clearRect", x, y, width, height)
}

func (wrapper *CanvasRenderingContext2D) LineTo(x float64, y float64) {
	(*js.Value)(wrapper).Call("lineTo", x, y)
}

func (wrapper *CanvasRenderingContext2D) MoveTo(x float64, y float64) {
	(*js.Value)(wrapper).Call("moveTo", x, y)
}

func (wrapper *CanvasRenderingContext2D) Arc(x float64, y float64, radius float64, startAngle float64, endAngle float64) {
	(*js.Value)(wrapper).Call("arc", x, y, radius, startAngle, endAngle)
}

func (wrapper *CanvasRenderingContext2D) FillText(text string, x float64, y float64) {
	(*js.Value)(wrapper).Call("fillText", text, x, y)
}

func (wrapper *CanvasRenderingContext2D) SetGlobalAlpha(value float64) {
	(*js.Value)(wrapper).Set("globalAlpha", value)
}

func (wrapper *CanvasRenderingContext2D) GetGlobalAlpha() float64 {
	return (*js.Value)(wrapper).Get("globalAlpha").Float()
}

func (wrapper *CanvasRenderingContext2D) SetFillStyle(value string) {
	(*js.Value)(wrapper).Set("fillStyle", value)
}

func (wrapper *CanvasRenderingContext2D) GetFillStyle() string {
	return (*js.Value)(wrapper).Get("fillStyle").String()
}

func (wrapper *CanvasRenderingContext2D) SetFont(value string) {
	(*js.Value)(wrapper).Set("font", value)
}

func (wrapper *CanvasRenderingContext2D) GetFont() string {
	return (*js.Value)(wrapper).Get("font").String()
}

func (wrapper *CanvasRenderingContext2D) SetTextAlign(value string) {
	(*js.Value)(wrapper).Set("textAlign", value)
}

func (wrapper *CanvasRenderingContext2D) GetTextAlign() string {
	return (*js.Value)(wrapper).Get("textAlign").String()
}
