package result

import (
	"fmt"
	"testing"
)

func TestHttp(t *testing.T) {
	var resp interface{}

	// add data
	//resp["code"] = 200
	//resp["msg"] = "OK"
	//resp["data"] = "data"
	resp = map[interface{}]interface{}{
		"code": "st",
		"msg":  "OK",
		"data": "data",
		"ok":   true,
		"Data": map[string]string{"a": "b"},
	}
	//get Data
	//fmt.Println(resp["Data"])
	if s, ok := resp.(string); ok {
		fmt.Printf("resp is %v and type is %T \n", s, s)
	}

	fmt.Println(resp)
}
