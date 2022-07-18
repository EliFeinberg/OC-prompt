package libs

import (
	"github.com/c-bata/go-prompt"
)

// Take from Kube-Prompt https://github.com/c-bata/kube-prompt/blob/main/kube/completer.go
func PortSuggest() []prompt.Suggest {
	return []prompt.Suggest{
		{Text: "8080", Description: ""},
	}
}
