package libs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/c-bata/go-prompt"
)

// func ParseCommands(path string) map[string]string

func ParseFiletoSuggest(path string) []prompt.Suggest {
	// Opens Json file to be parsed
	contents, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	// Defer close so fd is closed at termination of function
	defer contents.Close()

	// Set of Commands returned
	var Commands []prompt.Suggest

	// Format Json into Go accessible values then compile a list of
	var result map[string]map[string]string
	jsonval, err := ioutil.ReadAll(contents)
	if err != nil {
		fmt.Println(err)
	}
	json.Unmarshal(jsonval, &result)
	for k, v := range result {
		val := prompt.Suggest{Text: k, Description: v["Description"]}
		Commands = append(Commands, val)
	}
	return Commands
}

func WriteHelp(str string, cmd string) {

	f, _ := os.Create("source/Parse-Code/"+ cmd + ".txt")

	defer f.Close()

	n2, _ := f.WriteString(str)
	fmt.Printf("wrote %d bytes\n", n2)
}

func ParseFileForCommandList(path string) []string {
	// Opens Json file to be parsed
	contents, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	// Defer close so fd is closed at termination of function
	defer contents.Close()

	// Set of Commands returned
	var Commands []string

	// Format Json into Go accessible values then compile a list of
	var result map[string]map[string]string
	jsonval, err := ioutil.ReadAll(contents)
	if err != nil {
		fmt.Println(err)
	}
	json.Unmarshal(jsonval, &result)
	for k, _ := range result {
		Commands = append(Commands, k)
	}
	return Commands
}

// Taken from online resource
// https://play.golang.org/p/Qg_uv_inCek
// contains checks if a string is present in a slice
func StringInList(elem string, list []string) bool {
	for _, v := range list {
		if v == elem {
			return true
		}
	}
	return false
}

// Adapted from https://stackoverflow.com/questions/37334119/how-to-delete-an-element-from-a-slice-in-golang
func Remove(list []prompt.Suggest, index int) []prompt.Suggest {
	list[index] = list[len(list)-1]
	return list[:len(list)-1]
}
