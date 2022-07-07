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
var GlobalFlags = false
var prune = true
var helpOp = true
var stderr bytes.Buffer

func completer(d prompt.Document) []prompt.Suggest {
	CMDargs := strings.Split(d.Text, " ")
	var s []prompt.Suggest

	if libs.StringInList(CMDargs[0], OC_COMMANDS) {
		if CMDargs[0] == "api-versions" {
			return []prompt.Suggest{}
		} else if CMDargs[0] == "exit" {
			return []prompt.Suggest{}
		} else if CMDargs[0] == "ex" {
			return []prompt.Suggest{}
		} else if CMDargs[0] == "logout" {
			return []prompt.Suggest{}
		} else {
			s = resources.CommandFlags()[CMDargs[0]]
		}
		// Extra Settings for Customization
		if GlobalFlags {
			s = append(s, GLOBAL_OP...)
		}
		if prune {
			pruneUsedArgs(CMDargs, &s)
		}

	} else {
		s = OC_COMMANDS_SUGGEST
	}
	if helpOp {
		s = append(s, prompt.Suggest{Text: "--help", Description: "for more information about a given command"})
	}

	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}

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
	OC_COMMANDS = resources.CommandList()
	OC_COMMANDS_SUGGEST = resources.Commands()
	GLOBAL_OP = resources.GlobalOps()
	var ps *exec.Cmd

	fmt.Println("OpenShift Interactive Command Line Interface")
	for {
		t := prompt.Input(
			"oc ",
			completer,
			prompt.OptionTitle("RHOCP CLI"),
			prompt.OptionSelectedDescriptionTextColor(prompt.DarkGray))
		CMDargs := strings.Split(t, " ")
		if CMDargs[0] == "exit" {
			os.Exit(0)
		}
		if CMDargs[0] == "clear" {
			ps = exec.Command("clear")
		} else {
			ps = exec.Command("oc", CMDargs...)
			// fmt.Println(CMDargs)
		}
		ps.Stderr = &stderr
		res, err := ps.Output()

		if err != nil {
			fmt.Println(stderr.String())
			stderr.Reset()
		} else {
			fmt.Println(string(res))
		}
	}
}
