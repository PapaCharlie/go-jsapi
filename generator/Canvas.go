package main

var Canvas = WasmWrapper{
	Name: "Canvas",
	Properties: []Property{
		{
			Name: "width",
			Type: JsFloat,
		},
		{
			Name: "height",
			Type: JsFloat,
		},
	},
}
