package main

import (
	"fmt"
	routes "immmodi/framework/routes"
	"log"
	"net/http"
	"reflect"
	"strings"
)

var methodsStringArray []string
var methods map[string]reflect.Method
var routesArray []Route

func main() {
	GetAllRouterMethods(&methodsStringArray, &methods, &routesArray)

	log.Println("Starting server http://127.0.0.1:8000")
	http.HandleFunc("/", routerHandler)
	http.ListenAndServe(":8000", nil)
}

func routerHandler(w http.ResponseWriter, r *http.Request) {
	for _, route := range routesArray {
		if r.URL.Path == "/"+route.RouteName {
			response, err := route.routeHandler(w)
			if err != nil {
				w.Write(([]byte(err.Error())))
			} else {
				w.Write((response))
			}
			return
		}
	}

	http.NotFound(w, r)
}

type Route struct {
	// Method      string
	RouteName string
	// ContentType string
	// Handler     func() []byte
}

func (route *Route) routeHandler(w http.ResponseWriter) ([]byte, error) {
	response, contentType, err := GetResponseByRoute(&methodsStringArray, &methods, route.RouteName)
	w.Header().Add("Content-Type", contentType)
	if err != nil {
		return []byte(err.Error()), err
	}

	return response, nil
}

func GetAllRouterMethods(methodsStringArray *[]string, globalMethods *map[string]reflect.Method, routesArray *[]Route) {
	var routerMethod routes.RouterMethod
	structType := reflect.TypeOf(routerMethod)

	methods := make(map[string]reflect.Method)

	for i := 0; i < structType.NumMethod(); i++ {
		method := structType.Method(i)
		*methodsStringArray = append(*methodsStringArray, method.Name)
		(*routesArray) = append((*routesArray), Route{
			RouteName: strings.ToLower(strings.TrimRight(method.Name, "Respon")),
		})
		methods[method.Name] = method
	}
	(*globalMethods) = methods
}

func RunMethodByName(methods *map[string]reflect.Method, methodName string) ([]byte, string, error) {
	var routerMethod routes.RouterMethod
	structValue := reflect.ValueOf(routerMethod)

	if method, exists := (*methods)[methodName]; exists {
		output := method.Func.Call([]reflect.Value{structValue})
		return output[0].Bytes(), output[1].String(), nil
	} else {
		fmt.Printf("Method %s not found\n", methodName)
		return nil, "text/plain; charset=utf-8", fmt.Errorf("method %s not found", methodName)
	}

}

func GetResponseByRoute(methodsStringArray *[]string, methods *map[string]reflect.Method, routeName string) ([]byte, string, error) {
	routeHandler, err := GetRouteHandlerString(routeName, methodsStringArray)
	if err != nil {
		return []byte(err.Error()), "text/plain; charset=utf-8", err
	}
	response, contentType, err := RunMethodByName(methods, routeHandler)
	if err != nil {
		return nil, contentType, err
	}
	return response, contentType, nil
}

func GetRouteHandlerString(routeName string, methodsStringArray *[]string) (string, error) {
	for _, funcName := range *methodsStringArray {
		print(strings.TrimLeft(routeName, "/"))
		print(" --- ")
		println(funcName)
		if strings.Contains(strings.ToLower(funcName), strings.TrimLeft(routeName, "/")) {
			return funcName, nil
		}
	}

	return "", fmt.Errorf("this route: %s doesn't exist", routeName)
}
