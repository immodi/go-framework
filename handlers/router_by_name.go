package handlers

import (
	"fmt"
	"net/http"
	"reflect"
	"strings"
	"time"
)

type Route struct {
	RouteName string
}

type RouterMethod struct{}

type RouterMethodsResult struct {
	Methods            map[string]reflect.Method
	MethodsStringArray []string
	RoutesArray        []Route
}

func RouterByName(fn func(http.ResponseWriter, *http.Request, *RouterMethodsResult)) http.HandlerFunc {
	//makes the handler and returns it
	return func(w http.ResponseWriter, r *http.Request) {
		ch := make(chan RouterMethodsResult, 1)
		var methodsData RouterMethodsResult

		go GetAllRouterMethods(ch)

		select {
		case result := <-ch:
			methodsData = result
		case <-time.After(time.Second * 5):
			panic("Failed To get data")
		}
		fn(w, r, &methodsData)
	}
}

// the http route handler that will be called on '/'
func RootHandler(w http.ResponseWriter, r *http.Request, methodsData *RouterMethodsResult) {

	// loop through all RouteHanlder method routes
	for _, route := range methodsData.RoutesArray {
		var routeName string

		// if cuurent route is a route in those methods
		if r.URL.Path == route.RouteName {

			if route.RouteName == "/" {
				routeName = "root"
			} else {
				routeName = route.RouteName
			}

			// run the method and get response
			response, err := RouteHandler(w, r, &methodsData.MethodsStringArray, &methodsData.Methods, routeName)

			if err != nil {
				w.Write(([]byte(err.Error())))
			} else {
				w.Write(([]byte(response.Payload)))
			}

			return
		}
	}

	http.NotFound(w, r)
}

func GetAllRouterMethods(ch chan RouterMethodsResult) {
	// called asyncly at the start and returns all data related to methods that are binding to 'RouterMethod'

	var routerMethod RouterMethod
	structType := reflect.TypeOf(routerMethod)

	methods := make(map[string]reflect.Method)
	methodsStringArray := make([]string, 0)
	routesArray := make([]Route, 0)

	for i := 0; i < structType.NumMethod(); i++ {
		method := structType.Method(i)
		methodsStringArray = append(methodsStringArray, method.Name)

		var route string
		// if the method_name is 'root'
		if strings.ToLower(method.Name) == "root" {
			//change route to '/'
			route = "/"
		} else {
			// else change it to (/ + method_name)
			route = strings.ToLower("/" + method.Name)
		}

		(routesArray) = append((routesArray), Route{
			RouteName: route,
		})
		methods[method.Name] = method
	}

	result := RouterMethodsResult{
		Methods:            methods,
		MethodsStringArray: methodsStringArray,
		RoutesArray:        routesArray,
	}

	ch <- result
}

func RouteHandler(w http.ResponseWriter, r *http.Request, methodsStringArray *[]string, methods *map[string]reflect.Method, routeName string) (*Response, error) {
	response, err := GetResponseByRoute(methodsStringArray, methods, routeName, r)
	// http response to return

	w.Header().Add("Content-Type", response.ContentType)
	if err != nil {
		return response, err
	}

	return response, nil
}

func GetResponseByRoute(methodsStringArray *[]string, methods *map[string]reflect.Method, routeName string, r *http.Request) (*Response, error) {
	routeHandler, err := GetRouteHandlerString(routeName, methodsStringArray)
	response := &Response{}
	// the method name that will be called to get the response
	if err != nil {
		response.ContentType = "text/plain; charset=utf-8"
		response.Payload = []byte(err.Error())
		return response, err
	}

	response, err = RunMethodByName(methods, routeHandler, r)

	// run the method and return the response
	if err != nil {
		return response, err
	}

	return response, nil
}

func GetRouteHandlerString(routeName string, methodsStringArray *[]string) (string, error) {
	//checks if current route is in all router methods and returns method name
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

func RunMethodByName(methods *map[string]reflect.Method, methodName string, r *http.Request) (*Response, error) {
	var routerMethod RouterMethod
	structValue := reflect.ValueOf(routerMethod)
	requestStruct := reflect.ValueOf(r)

	if method, exists := (*methods)[methodName]; exists {
		output := method.Func.Call([]reflect.Value{structValue, requestStruct})
		response, ok := output[0].Interface().(*Response)
		if !ok {
			return &Response{}, fmt.Errorf("couldn't convert method: %s to route", methodName)
		}
		return response, nil

	} else {
		fmt.Printf("Method %s not found\n", methodName)
		return &Response{
			Payload:     []byte(fmt.Errorf("method %s not found", methodName).Error()),
			ContentType: "text/plain; charset=utf-8",
		}, fmt.Errorf("couldn't convert method: %s to route", methodName)
	}

}
