package main

var ArbitraryObject = WasmWrapper{
	Name: "ArbitraryObject",
	Properties: []Property{
		{
			Name: "innerHTML",
			Type: JsString,
		},
	},
}
