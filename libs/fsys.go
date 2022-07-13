package libs

import (
	"os"
	"strings"

	"github.com/c-bata/go-prompt/completer"
)

// Take from Kube-Prompt https://github.com/c-bata/kube-prompt/blob/main/kube/completer.go
var FileCompleter = completer.FilePathCompleter{
	IgnoreCase: true,
	Filter: func(fi os.FileInfo) bool {
		// if fi.IsDir() {
		// 	return true
		// }
		if strings.HasSuffix(fi.Name(), ".yaml") || strings.HasSuffix(fi.Name(), ".yml") {
			return true
		}
		return false
	},
}
