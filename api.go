package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"task_master/taskCatalogue"
)

func TaskMaster(w http.ResponseWriter, r *http.Request) {

	response, _ := ExecuteTask(IOStreamToByte(r.Body))
	fmt.Println(response)

	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func ListTasks(w http.ResponseWriter, _ *http.Request) {
	fmt.Println("listing")
	w.Header().Set("Content-Type", "application/json")
	tasks := taskCatalogue.TaskList()
	response := map[string]interface{}{
		"tasks": tasks,
	}

	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func IOStreamToByte(stream io.Reader) []byte {
	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(stream)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return buf.Bytes()
}
