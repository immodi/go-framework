package helpers

import (
	"encoding/json"
	"fmt"
)

func JsonStringParser(jsonString string) (*[]byte, string) {
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

	return &jsonData, "application/json"
}

func JsonStringMaker() string {
	return `{"test":"test","ajsdbssad":"j","asdkjkabshd":1,"bool":true,"aljsdbahsdba":[true,1,"asdahbds"]}`
}
