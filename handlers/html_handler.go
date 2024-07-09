package handlers

import (
	"bytes"
	"fmt"
	"html/template"
	"immmodi/framework/router"
	"net/http"
	"os"
	"reflect"
	"strconv"

	"github.com/a-h/templ"
)

func HtmlResponse(r *http.Request, template any, dataObject any) *router.Response {
	switch o := template.(type) {
	case string:
		return renderTemplate(o, dataObject)
	default:
		if component, isIt := isComponent(template, dataObject); isIt {
			return renderTmplComponent(r, component)
		}
		return &router.Response{
			ContentType: "text/plain",
			Payload:     []byte(fmt.Sprintf("Internal Server Error %s", strconv.Itoa(http.StatusInternalServerError))),
		}
	}
}

func renderTemplate(templateName string, dataObject any) *router.Response {
	if _, err := os.Stat("templates"); os.IsNotExist(err) {
		return &router.Response{
			ContentType: "text/plain",
			Payload:     []byte(fmt.Sprintln("Please Make Sure to put all the .html files inside the '/templates' directory")),
		}
	}

	tmpl, err := template.ParseFiles("templates/" + templateName + ".html")
	if err != nil {
		return &router.Response{
			ContentType: "text/plain",
			Payload:     []byte(fmt.Sprintf("Internal Server Error %s", strconv.Itoa(http.StatusInternalServerError))),
		}
	}

	var buf bytes.Buffer

	err = tmpl.Execute(&buf, dataObject)
	if err != nil {
		return &router.Response{
			ContentType: "text/plain",
			Payload:     []byte(fmt.Sprintf("Internal Server Error %s", strconv.Itoa(http.StatusInternalServerError))),
		}
	}

	renderedHTML := buf.String()

	return &router.Response{
		Payload:     []byte(renderedHTML),
		ContentType: "text/html",
	}
}

func renderTmplComponent(r *http.Request, component templ.Component) *router.Response {
	var buf bytes.Buffer

	component.Render(r.Context(), &buf)
	renderedHTML := buf.String()

	return &router.Response{
		Payload:     []byte(renderedHTML),
		ContentType: "text/html",
	}
}

func isComponent(template any, dataObject any) (templ.Component, bool) {
	tmplValue := reflect.ValueOf(template)
	tmplType := tmplValue.Type()

	if tmplType.Kind() == reflect.Func && tmplType.NumIn() == 1 && tmplType.NumOut() == 1 {
		if tmplType.Out(0) == reflect.TypeOf((*templ.Component)(nil)).Elem() {
			input := reflect.ValueOf(dataObject)
			result := tmplValue.Call([]reflect.Value{input})[0]
			component, ok := result.Interface().(templ.Component)
			if ok {
				return component, true
			}
		}
	}

	return nil, false
}
