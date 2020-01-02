package logutils

import (
	"encoding/json"
	"log"
)

//打印更漂亮的对象
func PrintObj(obj interface{}) {
	jsIndent, _ := json.MarshalIndent(&obj, "", "\t")

	log.Println(string(jsIndent))

}
