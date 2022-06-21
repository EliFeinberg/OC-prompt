package main

import (
	"fmt"
	"go-prompt/libs"
	"os"
	"os/exec"
	"strings"

	"github.com/c-bata/go-prompt"
)

func completer(d prompt.Document) []prompt.Suggest {
	// text := d.GetWordBeforeCursor()
	// fmt.Println(text)
	s := libs.ParseFiletoSuggest("source/commands.json")

	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}
func main() {
	fmt.Println("OpenShift Interactive Command Line Interface")
	for {
		t := prompt.Input(
			"oc ",
			completer,
			prompt.OptionTitle("RHOCP CLI"),
			prompt.OptionSelectedDescriptionTextColor(prompt.DarkGray))
		var ps *exec.Cmd
		CMDargs := strings.Split(t, " ")
		if CMDargs[0] == "exit" {
			os.Exit(0)
		}
		if CMDargs[0] == "clear" {
			ps = exec.Command("clear")
		} else {
			ps = exec.Command("oc", CMDargs...)
		}

		res, err := ps.Output()

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(string(res))
		}
	}
}
