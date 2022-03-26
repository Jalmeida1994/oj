package apps

import (
	"encoding/json"
	"log"
	"strings"
)

type EnvOutput struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type EnvApp struct {
	Version     string
	Description string
	Command     string
}

func NewEnv() *EnvApp {
	return &EnvApp{
		Version:     "1.0.0",
		Description: "Enviroment variables to Json",
		Command:     "env",
	}
}

func (app EnvApp) Parse(text string) string {
	var lines = strings.Split(text, "\n")
	var output = []EnvOutput{}
	for _, line := range lines {
		var pieces = strings.Split(line, "=")
		if len(pieces) == 2 {
			output = append(output, EnvOutput{Name: pieces[0], Value: pieces[1]})
		} else {
			output = append(output, EnvOutput{Name: pieces[0], Value: ""})
		}
	}
	jsonData, err := json.MarshalIndent(output, "", "  ")
	if err != nil {
		log.Println(err)
	}
	return string(jsonData)
}
