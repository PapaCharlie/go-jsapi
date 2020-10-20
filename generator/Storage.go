package main

import "github.com/dave/jennifer/jen"

var Storage = WasmWrapper{
	Name: "Storage",
	Methods: []Method{
		{
			Name: "clear",
		},
		{
			Name: "getItem",
			Args: []Arg{
				{
					Name: "keyName",
					Type: jen.String(),
				},
			},
			ReturnType: JsRaw,
		},
		{
			Name: "setItem",
			Args: []Arg{
				{
					Name: "keyName",
					Type: jen.String(),
				},
				{
					Name: "keyValue",
					Type: jen.String(),
				},
			},
		},
	},
}
