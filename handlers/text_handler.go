package handlers

import (
	"immmodi/framework/router"
)

func TextResponse(payload string) *router.Response {
	return &router.Response{
		ContentType: "text/plain; charset=utf-8",
		Payload:     []byte(payload),
	}
}
