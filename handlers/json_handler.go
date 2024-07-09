package handlers

import (
	"immmodi/framework/helpers"
	"immmodi/framework/router"
)

func JsonHandler(jsonString string) *router.Response {
	jsonData, contentType := helpers.JsonStringParser(jsonString)

	response := router.Response{
		Payload:     *jsonData,
		ContentType: contentType,
	}

	return &response

}
