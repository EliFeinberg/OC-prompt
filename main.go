package main

import (
	"fmt"

	"github.com/c-bata/go-prompt"
)

func completer(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "users", Description: "Store the username and age"},
		{Text: "articles", Description: "Store the article text posted by user"},
		{Text: "comments", Description: "Store the text commented to articles"},
	}
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
