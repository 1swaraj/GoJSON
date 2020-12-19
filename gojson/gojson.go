// Package gojson : The goal of this project is to enable creation, parsing, modification of JSON in Go without knowing the go structure
package gojson

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// JSONElement references a specific element within a wrapped structure.
type JSONElement struct {
	jsonInterface interface{}
}

// JSONData returns the value of the jsonElement
func (json *JSONElement) JSONData() interface{} {
	if json != nil {
		return json.jsonInterface
	}
	return nil
}

// Search performs the search with wildcard on nested json
func (json *JSONElement) Search(nested ...string) (*JSONElement, error) {
	jsonInterface := json.JSONData()
	for i := 0; i < len(nested); i++ {
		jsonPath := nested[i]
		if mmap, ok := jsonInterface.(map[string]interface{}); ok {
			jsonInterface, ok = mmap[jsonPath]
			if !ok {
				return nil, fmt.Errorf("Error in resolving path '%v' for key '%v'", i, jsonPath)
			}
		} else if marray, ok := jsonInterface.([]interface{}); ok {
			if jsonPath == "*" {
				tmp := []interface{}{}
				for _, val := range marray {
					if i+1 >= len(nested) {
						tmp = append(tmp, val)
						continue
					}
					value := &JSONElement{val}
					res, ok := value.Search(nested[i+1:]...)
					if ok == nil {
						tmp = append(tmp, res.JSONData())
					}
				}
				if len(tmp) == 0 {
					return nil, nil
				}
				return &JSONElement{tmp}, nil
			}
			index, err := strconv.Atoi(jsonPath)
			if err != nil {
				return nil, fmt.Errorf("Error in parsing path '%v' for key '%v' Error : %v", i, jsonPath, err)
			}
			if index < 0 {
				return nil, fmt.Errorf("Error in parsing path '%v' for key '%v' Error : %v ", i, jsonPath, err)
			}
			if len(marray) <= index {
				return nil, fmt.Errorf("Error in parsing path '%v' for key '%v' overflows Error : %v", i, jsonPath, err)
			}
			jsonInterface = marray[index]
		} else {
			return nil, fmt.Errorf("Error in parsing path '%v' for key '%v' ", i, jsonPath)
		}
	}
	return &JSONElement{jsonInterface}, nil
}

// ParseJSON unmarshals a JSON byte slice into a *JSONElement.
func ParseJSON(sample []byte) (*JSONElement, error) {
	var gojson JSONElement

	if err := json.Unmarshal(sample, &gojson.jsonInterface); err != nil {
		return nil, err
	}

	return &gojson, nil
}
