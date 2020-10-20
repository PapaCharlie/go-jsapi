package jsapi

import "syscall/js"

type D3Transform js.Value

func (wrapper *D3Transform) Apply(x, y float64) (float64, float64) {
	return (*js.Value)(wrapper).Call("applyX", x).Float(), (*js.Value)(wrapper).Call("applyY", y).Float()
}
