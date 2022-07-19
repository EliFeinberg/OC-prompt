package resources

import (
	"go-prompt/libs"
	"strings"

	"github.com/c-bata/go-prompt"
)

type UsageType func(prompt.Document, *[]prompt.Suggest)

var usage map[string]string = map[string]string{
	"new-build":       "(IMAGE | IMAGESTREAM | PATH | URL ...) [flags]",
	"adm":             "[flags]",
	"wait":            "([-f FILENAME] | resource.group/resource.name | resource.group [(-l label | --all)]) [--for=delete|--for condition=available|--for=jsonpath='{}'=value] [flags]",
	"scale":           "[--resource-version=version] [--current-replicas=count] --replicas=COUNT (-f FILENAME | TYPE NAME) [flags]",
	"policy":          "[flags]",
	"get":             "[(-o|--output=)json|yaml|name|go-template|go-template-file|template|templatefile|jsonpath|jsonpath-as-json|jsonpath-file|custom-columns-file|custom-columns|wide] (TYPE[.VERSION][.GROUP] [NAME | -l label] | TYPE[.VERSION][.GROUP]/NAME ...) [flags]",
	"import-image":    "IMAGESTREAM[:TAG] [flags]",
	"delete":          "([-f FILENAME] | [-k DIRECTORY] | TYPE [(NAME | -l label | --all)]) [flags]",
	"cancel-build":    "(BUILD | BUILDCONFIG) [flags]",
	"version":         "[flags]",
	"start-build":     "(BUILDCONFIG | --from-build=BUILD) [flags]",
	"diff":            "-f FILENAME [flags]",
	"set":             "COMMAND [flags]",
	"logout":          "[flags]",
	"registry":        "COMMAND [flags]",
	"port-forward":    "TYPE/NAME [options] [LOCAL_PORT:]REMOTE_PORT [...[LOCAL_PORT_N:]REMOTE_PORT_N] [flags]",
	"api-versions":    "[flags]",
	"exec":            "(POD | TYPE/NAME) [-c CONTAINER] [flags] -- COMMAND [args...]",
	"logs":            "[-f] [-p] (POD | TYPE/NAME) [-c CONTAINER] [flags]",
	"cp":              "<file-spec-src> <file-spec-dest> [flags]",
	"project":         "[NAME] [flags]",
	"debug":           "RESOURCE/NAME [ENV1=VAL1 ...] [-c CONTAINER] [flags] [-- COMMAND]",
	"proxy":           "[--port=PORT] [--www=static-dir] [--www-prefix=prefix] [--api-prefix=prefix] [flags]",
	"config":          "SUBCOMMAND [flags]",
	"serviceaccounts": "[flags]",
	"explain":         "RESOURCE [flags]",
	"rollout":         "SUBCOMMAND [flags]",
	"secrets":         "[flags]",
	"attach":          "(POD | TYPE/NAME) -c CONTAINER [flags]",
	"edit":            "(RESOURCE/NAME | -f FILENAME) [flags]",
	"whoami":          "[flags]",
	"rollback":        "(DEPLOYMENTCONFIG | DEPLOYMENT) [flags]",
	"idle":            "(SERVICE_ENDPOINTS... | -l label | --all | --resource-names-file FILENAME) [flags]",
	"plugin":          "[flags]",
	"projects":        "[flags]",
	"rsh":             "[-c CONTAINER] [flags] (POD | TYPE/NAME) COMMAND [args...]]",
	"ex":              "[flags]",
	"create":          "-f FILENAME [flags]",
	"replace":         "-f FILENAME [flags]",
	"observe":         "RESOURCE [-- COMMAND ...] [flags]",
	"cluster-info":    "[flags]",
	"describe":        "(-f FILENAME | TYPE [NAME_PREFIX | -l label] | TYPE/NAME) [flags]",
	"api-resources":   "[flags]",
	"apply":           "(-f FILENAME | -k DIRECTORY) [flags]",
	"new-app":         "(IMAGE | IMAGESTREAM | TEMPLATE | PATH | URL ...) [flags]",
	"rsync":           "SOURCE DESTINATION [flags]",
	"tag":             "[--source=SOURCETYPE] SOURCE DEST [DEST ...] [flags]",
	"annotate":        "[--overwrite] (-f FILENAME | TYPE NAME) KEY_1=VAL_1 ... KEY_N=VAL_N [--resource-version=version] [flags]",
	"auth":            "[flags]",
	"process":         "(TEMPLATE | -f FILENAME) [-p=KEY=VALUE] [flags]",
	"patch":           "(-f FILENAME | TYPE NAME) [-p PATCH|--patch-file FILE] [flags]",
	"label":           "[--overwrite] (-f FILENAME | TYPE NAME) KEY_1=VAL_1 ... KEY_N=VAL_N [--resource-version=version] [flags]",
	"extract":         "RESOURCE/NAME [--to=DIRECTORY] [--keys=KEY ...] [flags]",
	"image":           "COMMAND [flags]",
	"expose":          "(-f FILENAME | TYPE NAME) [--port=port] [--protocol=TCP|UDP|SCTP] [--target-port=number-or-name] [--name=name] [--external-ip=external-ip-of-service] [--type=type] [flags]",
	"run":             "NAME --image=image [--env='key=value'] [--port=port] [--dry-run=server|client] [--overrides=inline-json] [--command] -- [COMMAND] [args...] [flags]",
	"status":          "[-o dot | --suggest ] [flags]",
	"new-project":     "NAME [--display-name=DISPLAYNAME] [--description=DESCRIPTION] [flags]",
	"autoscale":       "(-f FILENAME | TYPE NAME | TYPE/NAME) [--min=MINPODS] --max=MAXPODS [--cpu-percent=CPU] [flags]",
	"kustomize":       "DIR [flags]",
}

func Usage() map[string]UsageType {
	return map[string]UsageType{
		"set": set,
	}
}

func set(d prompt.Document, sug *[]prompt.Suggest) {
	CMDargs := strings.Split(d.Text, " ")
	cmds := CommandFlags("set")
	if len(CMDargs) == 2 {
		for i := 1; i < len(cmds); i++ {
			for j := 0; j < len(*sug); j++ {

				if !libs.SuggestInList(cmds[i], (*sug)) {
					*sug = libs.Remove(*sug, j)

				}
			}
		}
	} else if len(CMDargs) >= 3 {
		for i := 0; i < len(cmds); i++ {
			for j := 0; j < len(*sug); j++ {

				if cmds[i].Text == (*sug)[j].Text {
					*sug = libs.Remove(*sug, j)

				}
			}
		}
	}
}
