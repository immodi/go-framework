package views

import (
	"immmodi/framework/components"
	"immmodi/framework/handlers"
	"immmodi/framework/helpers"
	"immmodi/framework/router"
	"net/http"
	"strconv"
)

func Test(r *http.Request) *router.Response {
	test := components.HtmlData{
		Title: "Ahmed",
	}

	return handlers.HtmlResponse(r, components.Test, test)
}

func Json(r *http.Request) *router.Response {
	j := helpers.JsonConstructor{}

	jsonString := j.JParseObject(
		"test",
		j.JParseArray(
			j.JParseBool(true),
			j.JParseInt(1414),
			j.JParseObject(`bdahbdas"`, j.JParseBool(false)),
		),
	).String()

	return handlers.JsonResonse(jsonString)
}

func Root(r *http.Request) *router.Response {
	return handlers.HtmlResponse(r, "index", nil)
}

func Text(r *http.Request) *router.Response {
	return handlers.TextResponse("a test repoinse man")
}

func BasicEdit(r *http.Request) *router.Response {
	integer := 0
	return handlers.HtmlResponse(r, components.BasicEdit, strconv.Itoa(integer))
}

func BasicEditPost(r *http.Request) *router.Response {
	numberString := r.FormValue("number")
	number, err := strconv.Atoi(numberString)
	if err != nil {
		panic(err.Error())
	}
	return handlers.HtmlResponse(r, components.BasicEdit, strconv.Itoa(number+1))
}
