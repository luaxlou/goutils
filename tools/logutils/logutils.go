package logutils

import (
	"encoding/json"
	"log"
)

//打印更漂亮的对象
func PrintObj(obj interface{}) {

	log.Println(FormatJSON(obj))

}

func FormatJSON(obj interface{}) string {
	jsIndent, _ := json.MarshalIndent(&obj, "", "\t")

	return string(jsIndent)
}
