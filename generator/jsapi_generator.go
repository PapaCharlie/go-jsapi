package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/dave/jennifer/jen"
)

type Arg struct {
	Name string
	Type *jen.Statement
}

type JsType struct {
	Type      *jen.Statement
	Transform *jen.Statement
}

var jsValue = jen.Qual("syscall/js", "Value")

var (
	JsRaw = &JsType{
		Type:      jsValue,
		Transform: jen.Empty(),
	}
	JsString = &JsType{
		Type:      jen.String(),
		Transform: jen.Dot("String").Call(),
	}
	JsInt = &JsType{
		Type:      jen.Int(),
		Transform: jen.Dot("Int").Call(),
	}
	JsFloat = &JsType{
		Type:      jen.Float64(),
		Transform: jen.Dot("Float").Call(),
	}
	JsBool = &JsType{
		Type:      jen.Bool(),
		Transform: jen.Dot("Bool").Call(),
	}
)

type Method struct {
	Name       string
	Doc        string
	Args       []Arg
	ReturnType *JsType
}

type Property struct {
	Name string
	Doc  string
	Type *JsType
}

type WasmWrapper struct {
	Name       string
	Doc        string
	Methods    []Method
	Properties []Property
}

var Wrappers = []WasmWrapper{
	Canvas,
	CanvasRenderingContext2D,
	Storage,
	ArbitraryObject,
}

func Export(s string) string {
	return strings.ToUpper(s[0:1]) + s[1:]
}

func (w *WasmWrapper) Generate() (err error) {
	code := jen.NewFile("jsapi")

	t := Export(w.Name)
	code.HeaderComment("+build js,wasm")
	code.PackageComment(`DO NOT EDIT

Code automatically generated by github.com/PapaCharlie/go-jsapi`)
	if w.Doc != "" {
		code.Comment(w.Doc)
	}
	code.Type().Id(t).Add(jsValue).Line()

	wrapper := jen.Code(jen.Id("wrapper"))
	for _, m := range w.Methods {
		if m.Doc != "" {
			code.Comment(w.Doc)
		}
		f := code.Func().Params(jen.Add(wrapper).Op("*").Id(t)).Id(Export(m.Name)).ParamsFunc(func(def *jen.Group) {
			for _, a := range m.Args {
				def.Id(a.Name).Add(a.Type)
			}
		})
		if m.ReturnType != nil {
			f.Add(m.ReturnType.Type)
		}
		f.BlockFunc(func(def *jen.Group) {
			c := jen.Call(jen.Op("*").Add(jsValue)).Call(wrapper).Dot("Call").CallFunc(func(def *jen.Group) {
				def.Lit(m.Name)
				for _, a := range m.Args {
					def.Id(a.Name)
				}
			})
			if m.ReturnType != nil {
				def.Return(c.Add(m.ReturnType.Transform))
			} else {
				def.Add(c)
			}
		}).Line()
	}

	value := jen.Code(jen.Id("value"))
	for _, p := range w.Properties {
		if p.Doc != "" {
			code.Comment(w.Doc)
		}
		code.Func().Params(jen.Add(wrapper).Op("*").Id(t)).Id("Set" + Export(p.Name)).
			Params(jen.Add(value).Add(p.Type.Type)).
			BlockFunc(func(def *jen.Group) {
				def.Call(jen.Op("*").Add(jsValue)).Call(wrapper).Dot("Set").Call(jen.Lit(p.Name), value)
			}).Line()
		if p.Doc != "" {
			code.Comment(w.Doc)
		}
		code.Func().Params(jen.Add(wrapper).Op("*").Id(t)).Id("Get" + Export(p.Name)).Params().Add(p.Type.Type).
			BlockFunc(func(def *jen.Group) {
				def.Return(jen.Call(jen.Op("*").Add(jsValue)).Call(wrapper).Dot("Get").Call(jen.Lit(p.Name)).Add(p.Type.Transform))
			}).Line()
	}

	target := w.Name + ".go"
	if err = os.Remove(target); err != nil && !os.IsNotExist(err) {
		return err
	}
	buf := &bytes.Buffer{}
	if err = code.Render(buf); err != nil {
		log.Println(err)
		return err
	}
	if err = ioutil.WriteFile(target, buf.Bytes(), 0444); err != nil {
		return err
	}
	return nil
}

func main() {
	for _, w := range Wrappers {
		if err := w.Generate(); err != nil {
			panic(err)
		}
	}
}
