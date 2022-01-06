package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"reflect"

	"task_master/taskCatalogue"
)

type CallObjTskData struct {
	FunctionName string                 `json:"function_name"`
	Parameters   map[string]interface{} `json:"parameters"`
}
type CallTskObj struct {
	Data     CallObjTskData `json:"data"`
	MetaData string         `json:"meta_data"`
}

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", TaskMaster)
	router.HandleFunc("/tasks", ListTasks)
	log.Fatal(http.ListenAndServe(":8080", router))

}

func ExecuteTask(request []byte) (result map[string]interface{}, err error) {
	var obj CallTskObj
	err = json.Unmarshal(request, &obj)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	_, found := taskCatalogue.TaskCatalogue[obj.Data.FunctionName]
	if !found {
		fmt.Println("Invalid  function  name")
		return nil, errors.New("Invalid  function  name")
	}
	f := reflect.ValueOf(taskCatalogue.TaskCatalogue[obj.Data.FunctionName])
	in := make([]reflect.Value, 1)
	in[0] = reflect.ValueOf(obj.Data.Parameters)

	var res []reflect.Value
	res = f.Call(in)
	if res[1].Interface() != nil {
		fmt.Println(err)
	}
	result = res[0].Interface().(map[string]interface{})
	return
}
