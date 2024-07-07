package routes

import (
	"fmt"
	"immmodi/framework/helpers"
	"immmodi/framework/types"
	"net/http"
)

type RouterMethod types.RouterMethod

func (RouterMethod) Test(r *http.Request) *types.Response {
	println(r.Method)
	response := types.Response{
		Payload:     []byte(fmt.Sprintf("<h1>%s</h1>", "sladhkasbdjashdbaj")),
		ContentType: "text/html",
	}

	return &response
}

func (RouterMethod) J(r *http.Request) *types.Response {
	println(r.Method)
	jsonString := helpers.JsonStringMaker()
	jsonData, contentType := helpers.JsonStringParser(jsonString)

	response := types.Response{
		Payload:     *jsonData,
		ContentType: contentType,
	}

	return &response
}

func (RouterMethod) Root(r *http.Request) *types.Response {
	response := types.Response{
		Payload:     []byte(fmt.Sprintf("<h1>%s</h1>", "Hello, World!")),
		ContentType: "text/html",
	}

	return &response
}
