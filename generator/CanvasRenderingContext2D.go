package main

import "github.com/dave/jennifer/jen"

var CanvasRenderingContext2D = WasmWrapper{
	Name: "CanvasRenderingContext2D",
	Methods: []Method{
		{
			Name: "beginPath",
		},
		{
			Name: "stroke",
		},
		{
			Name: "fill",
		},
		{
			Name: "fillRect",
			Args: []Arg{
				{
					Name: "x",
					Type: jen.Float64(),
				},
				{
					Name: "y",
					Type: jen.Float64(),
				},
				{
					Name: "width",
					Type: jen.Float64(),
				},
				{
					Name: "height",
					Type: jen.Float64(),
				},
			},
		},
		{
			Name: "strokeRect",
			Args: []Arg{
				{
					Name: "x",
					Type: jen.Float64(),
				},
				{
					Name: "y",
					Type: jen.Float64(),
				},
				{
					Name: "width",
					Type: jen.Float64(),
				},
				{
					Name: "height",
					Type: jen.Float64(),
				},
			},
		},
		{
			Name: "clearRect",
			Args: []Arg{
				{
					Name: "x",
					Type: jen.Float64(),
				},
				{
					Name: "y",
					Type: jen.Float64(),
				},
				{
					Name: "width",
					Type: jen.Float64(),
				},
				{
					Name: "height",
					Type: jen.Float64(),
				},
			},
		},
		{
			Name: "lineTo",
			Args: []Arg{
				{
					Name: "x",
					Type: jen.Float64(),
				},
				{
					Name: "y",
					Type: jen.Float64(),
				},
			},
		},
		{
			Name: "moveTo",
			Args: []Arg{
				{
					Name: "x",
					Type: jen.Float64(),
				},
				{
					Name: "y",
					Type: jen.Float64(),
				},
			},
		},
		{
			Name: "arc",
			Args: []Arg{
				{
					Name: "x",
					Type: jen.Float64(),
				},
				{
					Name: "y",
					Type: jen.Float64(),
				},
				{
					Name: "radius",
					Type: jen.Float64(),
				},
				{
					Name: "startAngle",
					Type: jen.Float64(),
				},
				{
					Name: "endAngle",
					Type: jen.Float64(),
				},
			},
		},
		{
			Name: "fillText",
			Args: []Arg{
				{
					Name: "text",
					Type: jen.String(),
				},
				{
					Name: "x",
					Type: jen.Float64(),
				},
				{
					Name: "y",
					Type: jen.Float64(),
				},
			},
		},
	},
	Properties: []Property{
		{
			Name: "globalAlpha",
			Type: JsFloat,
		},
		{
			Name: "fillStyle",
			Type: JsString,
		},
		{
			Name: "font",
			Type: JsString,
		},
		{
			Name: "textAlign",
			Type: JsString,
		},
	},
}
