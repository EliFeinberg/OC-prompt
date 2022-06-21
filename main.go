package main

import (
	"fmt"
	"go-prompt/libs"

	"github.com/c-bata/go-prompt"
)

func completer(d prompt.Document) []prompt.Suggest {
	text := d.GetWordBeforeCursor()
	// fmt.Println(text)
	s := libs.ParseFiletoSuggest("source/commands.json")

	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}
func main() {
	fmt.Println("OpenShift Interactive Command Line Interface")
	t := prompt.Input(
		"oc ",
		completer,
		prompt.OptionTitle("RHOCP CLI"),
		prompt.OptionSelectedDescriptionTextColor(prompt.DarkGray), prompt.OptionPrefixBackgroundColor(prompt.DarkBlue))

	fmt.Println("You selected: oc", t)
}
