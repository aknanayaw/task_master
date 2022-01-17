package taskCatalogue

import "task_master/task"

func TaskList() (tasks []Task) {
	var sayhi Task
	sayhi.Name = "say_hi"
	name := Params{
		Name: "name",
		Type: "string",
	}
	sayhi.Parameters = append(sayhi.Parameters, name)
	tasks = append(tasks, sayhi)

	var laptopgooglesearch Task
	laptopgooglesearch.Name = "laptop_google_search"
	searchString := Params{
		Name: "search_string",
		Type: "string",
	}
	laptopgooglesearch.Parameters = append(laptopgooglesearch.Parameters, searchString)
	tasks = append(tasks, laptopgooglesearch)
	return
}

var TaskCatalogue = map[string]interface{}{
	"say_hi": task.SayHi,

	"laptop_google_search": task.LaptopGoogleSearch}

type Task struct {
	Name       string   `json:"name"`
	Parameters []Params `json:"parameters"`
}

type Params struct {
	Name string `json:"name"`
	Type string `json:"type"`
}
