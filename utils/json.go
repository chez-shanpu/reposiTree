package utils

import (
	"encoding/json"
	"fmt"
	"github.com/chez-shanpu/reposiTree/pkg/tree"
	"io/ioutil"
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

func ReadJson(jsonName string, model *tree.NodeInfo) error {
	data, err := ioutil.ReadFile(jsonName)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, &model)
	if err != nil {
		return err
	}
	return nil
}
