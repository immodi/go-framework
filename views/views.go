package views

import (
	"immmodi/framework/handlers"
	"immmodi/framework/helpers"
	"immmodi/framework/router"
	"net/http"
)

type HtmlData struct {
	Title string
}

func Test(r *http.Request) *router.Response {
	test := HtmlData{
		Title: "Ahmed",
	}

	return handlers.HtmlResponse(r, hello, test)
}

func J(r *http.Request) *router.Response {
	j := helpers.JsonConstructor{}

	jsonString := j.JParseObject(
		"test",
		j.JParseArray(
			j.JParseBool(true),
			j.JParseInt(1414),
			j.JParseObject(`bdahbdas"`, j.JParseBool(false)),
		),
	).String()

	return handlers.JsonHandler(jsonString)
}

func Root(r *http.Request) *router.Response {
	return handlers.HtmlResponse(r, "index", nil)
}
