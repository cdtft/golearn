package json

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"
)

func TestMarshalIndentTest(t *testing.T) {

	c := make(map[string] interface{})
	c["name"] = "wangcheng"
	c["title"] = "golang"
	c["contact"] = map[string]interface{}{
		"home": "112",
		"cell": "755391",
	}

	//prefix每行的前缀 indent缩进字符
	data, err := json.MarshalIndent(c, "", "    ")
	if err != nil {
		log.Println("Error", err)
		return
	}
	fmt.Println(string(data))
}
