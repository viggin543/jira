package common

import (
	"encoding/json"
	"fmt"
	"github.com/yalp/jsonpath"
	"net/http"
	"os"
)

func ParseToSplitStr(body []byte, jsonPath string) [][2]string {
	parsed := JPathGet(body, jsonPath)
	return toSplitStr(parsed)
}

func toSplitStr(parsed interface{}) [][2]string {
	result := [][2]string{}
	var user [2]string
	for idx, v := range parsed.([]interface{}) {
		if idx % 2 == 0 {
			user = [2]string{}
			user[0] = v.(string)
		} else {
			user[1] = v.(string)
			result = append(result, user)
		}
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

func PanicIfNonEmpty(err error, response *http.Response) {
	if err != nil {
		fmt.Print(err,response)
		os.Exit(1)
	}
}


