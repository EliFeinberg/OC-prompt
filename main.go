package main

import (
	"bytes"
	"fmt"
	"go-prompt/libs"
	"go-prompt/resources"
	"os"
	"os/exec"
	"strings"

	"github.com/c-bata/go-prompt"
)

var OC_COMMANDS_SUGGEST []prompt.Suggest
var OC_COMMANDS []string
var GLOBAL_OP []prompt.Suggest
var GlobalFlags = true
var prune = true
var helpOp = false
var stderr bytes.Buffer
var histBuff []string

func completer(d prompt.Document) []prompt.Suggest {
	CMDargs := strings.Split(d.Text, " ")
	var s []prompt.Suggest
	prompt.OptionInputTextColor(prompt.Purple)
	// Check if command is valid
	if libs.StringInList(CMDargs[0], OC_COMMANDS) {

		// Commands without suggestions
		if CMDargs[0] == "api-versions" {
			return []prompt.Suggest{}
		} else if CMDargs[0] == "exit" {
			return []prompt.Suggest{}
		} else if CMDargs[0] == "ex" {
			return []prompt.Suggest{}
		} else if CMDargs[0] == "logout" {
			return []prompt.Suggest{}
		} else if len(CMDargs) > 2 && CMDargs[len(CMDargs)-2] == "-f" {
			// Completer for Local YAML files
			s = libs.FileCompleter.Complete(d)
		} else if strings.Contains(CMDargs[len(CMDargs)-1], "--port=") {
			// Completer for Local YAML files
			s = libs.PortSuggest()
		} else {
			// Suggestions for valid commands
			s = resources.CommandFlags(CMDargs[0])
		}
		// Extra Settings for Customization
		if GlobalFlags {
			s = append(s, GLOBAL_OP...)
		}
		if libs.StringInList(CMDargs[0], resources.Implemented) {
			resources.Usage()[CMDargs[0]](d, &s)
		}
		if prune {
			pruneUsedArgs(CMDargs, &s)
		}

	} else if len(CMDargs) < 2 {
		// List of all commands
		s = OC_COMMANDS_SUGGEST
	}
	if helpOp {
		// Option ofr help (might be already included in globalflags)
		s = append(s, prompt.Suggest{Text: "--help", Description: "for more information about a given command"})
	}

	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}

// Function for removing used args from suggestions as to not dupplicate
func pruneUsedArgs(CMDargs []string, s *[]prompt.Suggest) {
	for i := 1; i < len(CMDargs); i++ {
		for j := 0; j < len(*s); j++ {
			if CMDargs[i] == (*s)[j].Text {
				*s = libs.Remove(*s, j)
				break
			}
		}
	}
}
func main() {

	// Setting Up Global Variables
	OC_COMMANDS = resources.CommandList()
	OC_COMMANDS_SUGGEST = resources.Commands()
	GLOBAL_OP = resources.GlobalOps()
	var ps *exec.Cmd

	fmt.Println("OpenShift Interactive Command Line Interface")

	// Interface Loop for accepting Commands
	for {
		t := prompt.Input(
			"OC >>> ",
			completer,
			prompt.OptionTitle("RHOCP CLI"),
			prompt.OptionSuggestionTextColor(prompt.White),
			prompt.OptionSuggestionBGColor(prompt.DarkBlue),
			prompt.OptionSelectedDescriptionBGColor(prompt.Black),
			prompt.OptionPreviewSuggestionTextColor(prompt.Cyan),
			prompt.OptionSelectedSuggestionBGColor(prompt.Black),
			prompt.OptionSelectedDescriptionTextColor(prompt.White),
			prompt.OptionSelectedSuggestionTextColor(prompt.White),
			prompt.OptionSelectedDescriptionTextColor(prompt.White),
			prompt.OptionDescriptionTextColor(prompt.White),
			prompt.OptionDescriptionBGColor(prompt.Blue),
			prompt.OptionHistory(histBuff))

		// Seperate Arguments of Command
		CMDargs := strings.Split(t, " ")

		// Special case to exit
		if CMDargs[0] == "exit" {
			os.Exit(0)
		}
		// Add command to history
		histBuff = append(histBuff, t)

		// Special Case to clear screen
		if CMDargs[0] == "clear" {
			ps = exec.Command("clear")
		} else {
			// Actual Execution of Command
			if CMDargs[0] == "login" {
				log_us := os.Getenv("OC_USR")
				log_pass := os.Getenv("OC_PASS")
				log_port := os.Getenv("OC_PORT")
				if len(CMDargs) < 2 && log_port != "" {
					CMDargs = append(CMDargs, log_port)
					if log_us != "" {
						CMDargs = append(CMDargs, "-u", log_us)
						if log_pass != "" {
							CMDargs = append(CMDargs, "-p", log_pass)
						}
					}
				}
			}
			ps = exec.Command("oc", CMDargs...)
		}

		// Error Handling
		ps.Stderr = &stderr
		res, err := ps.Output()

		if err != nil {
			// Print and clear Standard Error
			fmt.Println(stderr.String())
			stderr.Reset()
		} else {
			// Print output of execution
			fmt.Println(string(res))
		}
	}
}
