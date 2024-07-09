package routes

import (
	"fmt"
	"immmodi/framework/handlers"
	"immmodi/framework/helpers"
	"net/http"
)

func Test(r *http.Request) *handlers.Response {
	return &handlers.Response{
		Payload:     []byte(fmt.Sprintf("<h1>%s</h1>", "sladhkasbdjashdbaj")),
		ContentType: "text/html",
	}
}

func J(r *http.Request) *handlers.Response {
	j := helpers.JsonConstructor{}

	jsonString := j.JParseObject(
		"test",
		j.JParseArray(
			j.JParseBool(true),
			j.JParseNil(),
		),
	).String()

	jsonData, contentType := helpers.JsonStringParser(jsonString)

	response := handlers.Response{
		Payload:     *jsonData,
		ContentType: contentType,
	}

	return &response
}

func Root(r *http.Request) *handlers.Response {
	return &handlers.Response{
		Payload:     []byte(fmt.Sprintf("<h1>%s</h1>", "Hello, World!")),
		ContentType: "text/html",
	}
}
