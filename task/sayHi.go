package task

import (
	"errors"
	"fmt"
)

func SayHi(inputObj map[string]interface{}) (map[string]interface{}, error) {
	name, inMap := inputObj["name"]
	if !inMap {
		fmt.Println("invalid parameters")
		return nil, errors.New("invalid parameters")
	}

	greeting := "Hi " + name.(string)
	fmt.Println(greeting)

	outputObj := map[string]interface{}{
		"greeting": greeting,
	}
	return outputObj, nil
}
