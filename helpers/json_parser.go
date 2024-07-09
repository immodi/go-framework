package helpers

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

type JsonConstructor struct {
	JsonString string
}

type JsonValue interface {
	jsonValue() string
}

type StringValue string
type IntValue string
type BoolValue string
type NilValue string
type ArrayValue string
type ObjectValue string

func (j JsonConstructor) JParseObject(key string, value JsonValue) ObjectValue {
	jsonString := "{" + string(j.JParseString(key)) + ":" + value.jsonValue() + "}"
	return ObjectValue(jsonString)
}

func (JsonConstructor) JParseString(jString string) StringValue {
	return StringValue(`"` + strings.ReplaceAll(jString, `"`, "") + `"`)
}

func (JsonConstructor) JParseInt(jInt int) IntValue {
	return IntValue(strconv.Itoa(jInt))
}

func (JsonConstructor) JParseBool(jBool bool) BoolValue {
	return BoolValue(strconv.FormatBool(jBool))
}

func (JsonConstructor) JParseNil() NilValue {
	return NilValue("null")
}

func (j JsonConstructor) JParseArray(values ...JsonValue) ArrayValue {
	jsonString := "["
	comma := ","

	for index, value := range values {
		if index == len(values)-1 {
			comma = ""
		}

		jsonString += value.jsonValue() + comma
	}
	return ArrayValue(jsonString + "]")
}

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

func (s StringValue) jsonValue() string { return s.String() }
func (s IntValue) jsonValue() string    { return s.String() }
func (s BoolValue) jsonValue() string   { return s.String() }
func (s NilValue) jsonValue() string    { return s.String() }
func (s ArrayValue) jsonValue() string  { return s.String() }
func (s ObjectValue) jsonValue() string { return s.String() }

func (v StringValue) String() string {
	return string(v)
}

func (v IntValue) String() string {
	return string(v)
}

func (v BoolValue) String() string {
	return string(v)
}

func (v NilValue) String() string {
	return string(v)
}

func (v ArrayValue) String() string {
	return string(v)
}

func (v ObjectValue) String() string {
	return string(v)
}
