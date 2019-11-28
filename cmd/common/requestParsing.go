package common

import (
	"encoding/json"
	"fmt"
	"github.com/yalp/jsonpath"
)

func ParseToSplitStr(body []byte, jsonPath string) []string {
	parsed := JPathGet(body, jsonPath)
	return toSplitStr(parsed)
}

func toSplitStr(parsed interface{}) []string {
	result := []string{}
	for _, v := range parsed.([]interface{}) {
		result = append(result, v.(string))
	}
	return result
}



func ParseToSting(body []byte, jsonPath string) string {
	parsed := JPathGet(body, jsonPath)
	return parsed.(string)
}


func JPathGet(body []byte, jsonPath string) interface{} {
	var response interface{}
	err := json.Unmarshal(body, &response)
	PanicIfNonEmpty(err, nil)
	parsed, err := jsonpath.Read(response, jsonPath)
	if err != nil {
		fmt.Println(response)
		panic(err)
	}
	return parsed
}




