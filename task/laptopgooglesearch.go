package task

import (
	"errors"
	"fmt"
	"github.com/pkg/browser"
	"github.com/rocketlaunchr/google-search"
)

func LaptopGoogleSearch(inputObj map[string]interface{}) (map[string]interface{}, error) {

	searchString, inMap := inputObj["search_string"]
	if !inMap {
		fmt.Println("invalid parameters")
		return nil, errors.New("invalid parameters")
	}

	result, err := googlesearch.Search(nil, searchString.(string))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	for i := 0; i < 3; i++ {
		err := browser.OpenURL(result[i].URL)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
	}

	outputObj := map[string]interface{}{
		"res1": result[0].Title,
		"res2": result[1].Title,
		"res3": result[2].Title,
	}
	return outputObj, nil
}
