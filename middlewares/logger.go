package middlewares

import (
	"fmt"
	"net/http"
	"time"
)

func Logger(r *http.Request) {
	formattedTime := time.Now().Format("Mon, 02 Jan 2006 15:04:05 MST")
	fmt.Printf("%s - %s - %s \n", r.URL, r.Method, formattedTime)
}
