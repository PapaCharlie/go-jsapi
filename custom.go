package jsapi

import "syscall/js"

func (wrapper *Canvas) Get2DContext() *CanvasRenderingContext2D {
	v := CanvasRenderingContext2D((*js.Value)(wrapper).Call("getContext", "2d"))
	return &v
}
