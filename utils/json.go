package utils

import (
	"encoding/json"
	"fmt"
	"os"
)

func OutputJson(outputName string, model interface{}) {
	file, err := os.Create(outputName)
	if err != nil {
		fmt.Print(err)
	}
	defer file.Close()
	bytes, _ := json.Marshal(model)
	if _, err := file.Write(bytes); err != nil {
		fmt.Print(err)
	}
}
