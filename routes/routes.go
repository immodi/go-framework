package routes

import (
	"encoding/json"
	"fmt"
)

type RouterMethod struct{}

func (RouterMethod) TestResponse() ([]byte, string) {
	return []byte(fmt.Sprintf("<h1>%s</h1>", "sladhkasbdjashdbaj")), "text/html"
}

func (RouterMethod) JResponse() ([]byte, string) {
	jsonString := `{"test":"test","ajsdbssad":"j","asdkjkabshd":1,"bool":true,"aljsdbahsdba":[true,1,"asdahbds"]}`

	var data map[string]any

	// Unmarshal JSON string into the map
	err := json.Unmarshal([]byte(jsonString), &data)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		panic(err)
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		panic(err)
	}

	return jsonData, "application/json"
}
