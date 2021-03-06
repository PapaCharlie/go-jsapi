// +build js,wasm

/*
DO NOT EDIT

Code automatically generated by github.com/PapaCharlie/go-jsapi
*/
package jsapi

import js "syscall/js"

type Storage js.Value

func (wrapper *Storage) Clear() {
	(*js.Value)(wrapper).Call("clear")
}

func (wrapper *Storage) GetItem(keyName string) js.Value {
	return (*js.Value)(wrapper).Call("getItem", keyName)
}

func (wrapper *Storage) SetItem(keyName string, keyValue string) {
	(*js.Value)(wrapper).Call("setItem", keyName, keyValue)
}
