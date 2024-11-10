package prettyprint

import (
	"encoding/json"
	"log"
)

func Log(data interface{}) {
	if prettyOutput, err := json.MarshalIndent(data, "", "  "); err == nil {
		log.Println(string(prettyOutput))
	} else {
		log.Println("don`t convert data to pretty-printed")
		log.Printf("%#v", data)
	}
}