package apps

import (
	"encoding/json"
	"fmt"
	"log"
	"regexp"
	"strings"
)

type PSOutput struct {
	Pid  string `json:"pid"`
	User string `json:"user"`
	Cpu  string `json:"cpu"`
	Mem  string `json:"mem"`
	Cmd  string `json:"cmd"`
}

type PSApp struct {
	Version     string
	Description string
	Command     string
}

func NewPS() *PSApp {
	return &PSApp{
		Version:     "1.0.0",
		Description: "Running processes to Json",
		Command:     "ps",
	}
}

func (app PSApp) Parse(text string) string {
	var lines = strings.Split(text, "\n")
	var output = []PSOutput{}
	for _, line := range lines {
		re := regexp.MustCompile(` +`)
		var pieces = re.Split(line, -1)
		fmt.Println(pieces)
		if len(pieces) == 4 {
			output = append(output, PSOutput{Pid: pieces[0], Cmd: pieces[3]})
		} else if len(pieces) == 11 {
			output = append(output, PSOutput{Pid: pieces[1], User: pieces[0], Cpu: pieces[2], Mem: pieces[3], Cmd: pieces[10]})
		} else {
			output = append(output, PSOutput{Pid: pieces[0]})
		}
	}
	jsonData, err := json.MarshalIndent(output, "", "  ")
	if err != nil {
		log.Println(err)
	}
	return string(jsonData)
}
