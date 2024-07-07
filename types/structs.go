package types

import (
	"reflect"
)

type RouterMethod struct{}

type Route struct {
	RouteName string
}

type RouterMethodsResult struct {
	Methods            map[string]reflect.Method
	MethodsStringArray []string
	RoutesArray        []Route
}

type Response struct {
	Payload     []byte
	ContentType string
}
